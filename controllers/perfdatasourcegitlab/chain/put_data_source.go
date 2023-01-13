package chain

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	perfApi "github.com/epam/edp-perf-operator/v2/api/edp/v1"
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
	gitLabSecretName = "gitlab-admin-password"
	keyName          = "name"
)

func (h PutDataSource) ServeRequest(dataSource *perfApi.PerfDataSourceGitLab) error {
	log.Info("start creating/updating GitLab data source in PERF", keyName, dataSource.Name)

	if err := h.tryToPutDataSource(dataSource); err != nil {
		setFailedStatus(dataSource)

		return err
	}

	setSuccessStatus(dataSource)

	log.Info("PERF DataSourceGitLab has been created.", keyName, dataSource.Name)

	return nil
}

func setFailedStatus(ds *perfApi.PerfDataSourceGitLab) {
	ds.Status.Status = "error"
}

func setSuccessStatus(ds *perfApi.PerfDataSourceGitLab) {
	ds.Status.Status = "created"
}

func (h PutDataSource) tryToPutDataSource(dsResource *perfApi.PerfDataSourceGitLab) error {
	ps, err := cluster.GetPerfServerCr(h.client, dsResource.Spec.PerfServerName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get perf server: %w", err)
	}

	dsReq, err := h.perfClient.GetProjectDataSource(ps.Spec.ProjectName, dsResource.Spec.Type)
	if err != nil {
		return fmt.Errorf("failed to get project data source: %w", err)
	}

	if dsReq != nil {
		log.Info("PERF GitLab data source already exists. try to update.", "type", dsResource.Spec.Type)

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

func (h PutDataSource) tryToUpdateDataSource(dsResource *perfApi.PerfDataSourceGitLab, dsReq *dto.DataSource) error {
	branchDiff := getBranchConfigDifference(dsResource, dsReq)
	repoDiff := getRepositoryConfigDifference(dsResource, dsReq)

	if branchDiff == nil && repoDiff == nil {
		log.Info("nothing to update in GitLab data source", keyName, dsReq.Name)

		return nil
	}

	s, err := cluster.GetSecret(h.client, gitLabSecretName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	dsCommand := command.GetGitLabDsUpdateCommand(dsReq, &command.DataSourceGitLabConfigDto{
		Type:         dsReq.Type,
		ApiUrl:       dsResource.Spec.Config.Url,
		Username:     string(s.Data["username"]),
		Password:     string(s.Data["password"]),
		Repositories: repoDiff,
		Branches:     branchDiff,
	})

	if err = h.perfClient.UpdateDataSource(dsCommand); err != nil {
		return fmt.Errorf("failed to update data source: %w", err)
	}

	return nil
}

func getBranchConfigDifference(dsResource *perfApi.PerfDataSourceGitLab, dsReq *dto.DataSource) []string {
	conf := common.ConvertToStringArray(dsReq.Config["branches"])

	return datasource.GetMissingElementsInDataSource(dsResource.Spec.Config.Branches, conf)
}

func getRepositoryConfigDifference(dsResource *perfApi.PerfDataSourceGitLab, dsReq *dto.DataSource) []string {
	conf := common.ConvertToStringArray(dsReq.Config["repositories"])

	return datasource.GetMissingElementsInDataSource(dsResource.Spec.Config.Repositories, conf)
}

func (h PutDataSource) createDataSource(projectName string, dsResource *perfApi.PerfDataSourceGitLab) error {
	s, err := cluster.GetSecret(h.client, gitLabSecretName, dsResource.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	dsCommand := command.GetGitLabDsCreateCommand(dsResource, string(s.Data["username"]), string(s.Data["password"]))
	if err = h.perfClient.CreateDataSource(projectName, dsCommand); err != nil {
		return fmt.Errorf("failed to create data source: %w", err)
	}

	return nil
}
