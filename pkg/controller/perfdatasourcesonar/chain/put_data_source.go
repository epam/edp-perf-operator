package chain

import (
	"github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
	"github.com/epam/edp-perf-operator/v2/pkg/util/datasource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PutDataSource struct {
	client     client.Client
	perfClient perf.PerfClient
}

const (
	sonarDataSourceSecretName = "sonar-admin-password"
)

func (h PutDataSource) ServeRequest(dataSource *v1alpha1.PerfDataSourceSonar) error {
	log.Info("start creating/updating Sonar data source in PERF", "name", dataSource.Name)
	if err := h.tryToPutDataSource(dataSource); err != nil {
		setFailedStatus(dataSource)
		return err
	}
	setSuccessStatus(dataSource)
	log.Info("PERF DataSourceSonar has been created.", "name", dataSource.Name)
	return nil
}

func setFailedStatus(ds *v1alpha1.PerfDataSourceSonar) {
	ds.Status.Status = "error"
}

func setSuccessStatus(ds *v1alpha1.PerfDataSourceSonar) {
	ds.Status.Status = "created"
}

func (h PutDataSource) tryToPutDataSource(dsResource *v1alpha1.PerfDataSourceSonar) error {
	ps, err := cluster.GetPerfServerCr(h.client, dsResource.Spec.PerfServerName, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsReq, err := h.perfClient.GetProjectDataSource(ps.Spec.ProjectName, dsResource.Spec.Type)
	if err != nil {
		return err
	}

	if dsReq != nil {
		log.Info("PERF Sonar data source already exists. try to update.", "type", dsResource.Spec.Type)
		if err := h.tryToActivateDataSource(dsReq, ps); err != nil {
			return err
		}
		return h.tryToUpdateDataSource(dsResource, dsReq)
	}

	return h.createDataSource(ps.Spec.ProjectName, dsResource)
}

func (h PutDataSource) tryToActivateDataSource(dsReq *dto.DataSource, ps *v1alpha1.PerfServer) error {
	if dsReq.Active {
		log.Info("PERF data source is already activated.", "name", dsReq.Name)
		return nil
	}
	return h.perfClient.ActivateDataSource(ps.Spec.ProjectName, dsReq.Id)
}

func (h PutDataSource) tryToUpdateDataSource(dsResource *v1alpha1.PerfDataSourceSonar, dsReq *dto.DataSource) error {
	diff := getConfigDifference(dsResource, dsReq)
	if len(diff) == 0 {
		log.Info("nothing to update in Sonar data source", "name", dsReq.Name)
		return nil
	}

	s, err := cluster.GetSecret(h.client, sonarDataSourceSecretName, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsCommand := command.GetSonarDsUpdateCommand(dsReq, command.DataSourceConfigDto{
		Type:       dsReq.Type,
		ApiUrl:     dsResource.Spec.Config.Url,
		Username:   string(s.Data["username"]),
		Password:   string(s.Data["password"]),
		Parameters: diff,
	})
	return h.perfClient.UpdateDataSource(dsCommand)
}

func getConfigDifference(dsResource *v1alpha1.PerfDataSourceSonar, dsReq *dto.DataSource) []string {
	conf := common.ConvertToStringArray(dsReq.Config["projectKeys"])
	return datasource.GetMissingElementsInDataSource(dsResource.Spec.Config.ProjectKeys, conf)
}

func (h PutDataSource) createDataSource(projectName string, dsResource *v1alpha1.PerfDataSourceSonar) error {
	s, err := cluster.GetSecret(h.client, sonarDataSourceSecretName, dsResource.Namespace)
	if err != nil {
		return err
	}

	dsCommand := command.GetSonarDsCreateCommand(dsResource, string(s.Data["username"]), string(s.Data["password"]))
	return h.perfClient.CreateDataSource(projectName, dsCommand)
}
