package chain

import (
	"encoding/json"
	"github.com/epmd-edp/perf-operator/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/pkg/client/perf"
	"github.com/epmd-edp/perf-operator/pkg/controller/perfdatasource/chain/handler"
	"github.com/epmd-edp/perf-operator/pkg/model/command"
	"github.com/epmd-edp/perf-operator/pkg/model/dto"
	"github.com/epmd-edp/perf-operator/pkg/util/cluster"
	"github.com/epmd-edp/perf-operator/pkg/util/consts"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PutDataSource struct {
	next       handler.PerfDataSourceHandler
	client     client.Client
	perfClient perf.PerfClient
}

func (h PutDataSource) ServeRequest(dataSource *v1alpha1.PerfDataSource) error {
	log.Info("start creating/updating data source in PERF", "name", dataSource.Name)
	if err := h.tryToPutDataSource(dataSource); err != nil {
		return err
	}
	log.Info("PERF DataSource has been created.", "name", dataSource.Name)
	return nil
}

func (h PutDataSource) tryToPutDataSource(dsResource *v1alpha1.PerfDataSource) error {
	ow := cluster.GetOwnerReference(consts.PerfServerKind, dsResource.GetOwnerReferences())
	ps, err := cluster.GetPerfServerCr(h.client, ow.Name, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsReq, err := h.perfClient.GetProjectDataSource(ps.Spec.ProjectName, dsResource.Spec.Name)
	if err != nil {
		return err
	}

	if dsReq == nil {
		err := h.perfClient.CreateDataSource(ps.Spec.ProjectName, command.ConvertToDataSourceCreateCommandModel(dsResource))
		if err != nil {
			return err
		}
		return nil
	}

	if err := h.tryToUpdateDataSource(dsReq, ps.Spec.ProjectName, dsResource); err != nil {
		return err
	}

	return nil
}

func (h PutDataSource) tryToUpdateDataSource(dsReq *dto.DataSource, projectName string, dsResource *v1alpha1.PerfDataSource) error {
	if err := h.activateDataSource(projectName, dsReq.Id, dsReq.Active); err != nil {
		return err
	}
	return h.updateDataSource(dsReq, dsResource)
}

func (h PutDataSource) activateDataSource(projectName string, dsId int, active bool) error {
	if !active {
		return h.perfClient.ActivateDataSource(projectName, dsId)
	}
	log.Info("datasource is already activated. skip activating...", "datasource id", dsId)
	return nil
}

func (h PutDataSource) updateDataSource(dsReq *dto.DataSource, dsResource *v1alpha1.PerfDataSource) error {
	var dsc map[string][]string
	if err := json.Unmarshal([]byte(dsResource.Spec.Config), &dsc); err != nil {
		return err
	}

	for k, v := range dsc {
		addToDsConfig(dsReq, k, v)
	}

	return h.perfClient.UpdateDataSource(*dsReq)
}

func addToDsConfig(ds *dto.DataSource, configKey string, configs interface{}) {
	slice := reflect.ValueOf(ds.Config[configKey])
	sliceNewConfig := reflect.ValueOf(configs)
	arr := make([]interface{}, slice.Len()+sliceNewConfig.Len())

	for i := 0; i < slice.Len(); i++ {
		arr[i] = slice.Index(i).Interface()
	}

	for i := 0; i < sliceNewConfig.Len(); i++ {
		arr[i+slice.Len()] = sliceNewConfig.Index(i).Interface()
	}
	ds.Config[configKey] = arr
}
