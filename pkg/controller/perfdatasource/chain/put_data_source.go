package chain

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/client/perf"
	"github.com/epmd-edp/perf-operator/v2/pkg/controller/perfdatasource/chain/handler"
	"github.com/epmd-edp/perf-operator/v2/pkg/model/command"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/cluster"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/consts"
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
		setFailedStatus(dataSource)
		return err
	}
	setSuccessStatus(dataSource)
	log.Info("PERF DataSource has been created.", "name", dataSource.Name)
	return nil
}

func setFailedStatus(ds *v1alpha1.PerfDataSource) {
	ds.Status.Status = "error"
}

func setSuccessStatus(ds *v1alpha1.PerfDataSource) {
	ds.Status.Status = "created"
}

func (h PutDataSource) tryToPutDataSource(dsResource *v1alpha1.PerfDataSource) error {
	ow := cluster.GetOwnerReference(consts.PerfServerKind, dsResource.GetOwnerReferences())
	ps, err := cluster.GetPerfServerCr(h.client, ow.Name, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsReq, err := h.perfClient.GetProjectDataSource(ps.Spec.ProjectName, dsResource.Spec.Type)
	if err != nil {
		return err
	}

	if dsReq != nil {
		log.Info("datasource already exists. skip creating.", "type", dsResource.Spec.Type)
		return nil
	}

	return h.createDataSource(ps.Spec.ProjectName, dsResource)
}

func (h PutDataSource) createDataSource(projectName string, dsResource *v1alpha1.PerfDataSource) error {
	s, err := cluster.GetSecret(h.client, dsResource.Spec.Config.CredentialName, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsCommand := command.ConvertToDataSourceCreateCommand(dsResource,
		string(s.Data["username"]), string(s.Data["password"]))
	return h.perfClient.CreateDataSource(projectName, dsCommand)
}
