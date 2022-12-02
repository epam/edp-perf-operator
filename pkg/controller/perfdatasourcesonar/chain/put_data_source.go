package chain

import (
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/pkg/apis/edp/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf"
	"github.com/epam/edp-perf-operator/v2/pkg/model/command"
	"github.com/epam/edp-perf-operator/v2/pkg/model/dto"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
	"github.com/epam/edp-perf-operator/v2/pkg/util/common"
	"github.com/epam/edp-perf-operator/v2/pkg/util/datasource"
)

type PutDataSource struct {
	client     client.Client
	perfClient perf.PerfClient
}

const (
	sonarDataSourceSecretName = "sonar-admin-password"
	keyName                   = "name"
)

func (h PutDataSource) ServeRequest(dataSource *perfApi.PerfDataSourceSonar) error {
	log.Info("start creating/updating Sonar data source in PERF", keyName, dataSource.Name)

	if err := h.tryToPutDataSource(dataSource); err != nil {
		setFailedStatus(dataSource)

		return err
	}

	setSuccessStatus(dataSource)

	log.Info("PERF DataSourceSonar has been created.", keyName, dataSource.Name)

	return nil
}

func setFailedStatus(ds *perfApi.PerfDataSourceSonar) {
	ds.Status.Status = "error"
}

func setSuccessStatus(ds *perfApi.PerfDataSourceSonar) {
	ds.Status.Status = "created"
}

func (h PutDataSource) tryToPutDataSource(dsResource *perfApi.PerfDataSourceSonar) error {
	ps, err := cluster.GetPerfServerCr(h.client, dsResource.Spec.PerfServerName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get perf server: %w", err)
	}

	dsReq, err := h.perfClient.GetProjectDataSource(ps.Spec.ProjectName, dsResource.Spec.Type)
	if err != nil {
		return fmt.Errorf("failed to get project data source: %w", err)
	}

	if dsReq != nil {
		log.Info("PERF Sonar data source already exists. try to update.", "type", dsResource.Spec.Type)

		if activateErr := h.tryToActivateDataSource(dsReq, ps); activateErr != nil {
			return activateErr
		}

		return h.tryToUpdateDataSource(dsResource, dsReq)
	}

	return h.createDataSource(ps.Spec.ProjectName, dsResource)
}

func (h PutDataSource) tryToActivateDataSource(dsReq *dto.DataSource, ps *perfApi.PerfServer) error {
	if dsReq.Active {
		log.Info("PERF data source is already activated.", keyName, dsReq.Name)

		return nil
	}

	if err := h.perfClient.ActivateDataSource(ps.Spec.ProjectName, dsReq.Id); err != nil {
		return fmt.Errorf("failed to activate data source: %w", err)
	}

	return nil
}

func (h PutDataSource) tryToUpdateDataSource(dsResource *perfApi.PerfDataSourceSonar, dsReq *dto.DataSource) error {
	diff := getConfigDifference(dsResource, dsReq)
	if len(diff) == 0 {
		log.Info("nothing to update in Sonar data source", keyName, dsReq.Name)

		return nil
	}

	s, err := cluster.GetSecret(h.client, sonarDataSourceSecretName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	dsCommand := command.GetSonarDsUpdateCommand(dsReq, &command.DataSourceConfigDto{
		Type:       dsReq.Type,
		ApiUrl:     dsResource.Spec.Config.Url,
		Username:   string(s.Data["username"]),
		Password:   string(s.Data["password"]),
		Parameters: diff,
	})

	if err = h.perfClient.UpdateDataSource(dsCommand); err != nil {
		return fmt.Errorf("failed to update data source: %w", err)
	}

	return nil
}

func getConfigDifference(dsResource *perfApi.PerfDataSourceSonar, dsReq *dto.DataSource) []string {
	conf := common.ConvertToStringArray(dsReq.Config["projectKeys"])

	return datasource.GetMissingElementsInDataSource(dsResource.Spec.Config.ProjectKeys, conf)
}

func (h PutDataSource) createDataSource(projectName string, dsResource *perfApi.PerfDataSourceSonar) error {
	s, err := cluster.GetSecret(h.client, sonarDataSourceSecretName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	dsCommand := command.GetSonarDsCreateCommand(dsResource, string(s.Data["username"]), string(s.Data["password"]))

	if err = h.perfClient.CreateDataSource(projectName, dsCommand); err != nil {
		return fmt.Errorf("failed to create data source: %w", err)
	}

	return nil
}
