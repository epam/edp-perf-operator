package command

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/model/dto"
	"strings"
)

type DataSourceType string

const (
	Jenkins DataSourceType = "JENKINS"
)

type DataSourceCommand struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Type   DataSourceType `json:"type"`
	Config interface{}    `json:"config"`
}

type DataSourceJenkinsConfig struct {
	JobNames []string `json:"jobNames"`
	Url      string   `json:"url"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

type DataSourceSonarConfig struct {
	ProjectKeys []string `json:"projectKeys"`
	Url         string   `json:"url"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
}

func ConvertToDataSourceCreateCommand(ds *v1alpha1.PerfDataSource, username, password string) DataSourceCommand {
	if Jenkins == DataSourceType(strings.ToUpper(ds.Spec.Type)) {
		return getJenkinsDsCreateCommand(ds, username, password)
	}
	return getSonarDsCreateCommand(ds, username, password)
}

func ConvertToDataSourceUpdateCommand(ds *v1alpha1.PerfDataSource, dsReq dto.DataSource) DataSourceCommand {
	if Jenkins == DataSourceType(ds.Spec.Type) {
		return getJenkinsDsUpdateCommand(ds, dsReq)
	}
	return getSonarDsUpdateCommand(ds, dsReq)
}

func getSonarDsCreateCommand(ds *v1alpha1.PerfDataSource, username string, password string) DataSourceCommand {
	return DataSourceCommand{
		Name: ds.Spec.Name,
		Type: DataSourceType(strings.ToUpper(ds.Spec.Type)),
		Config: DataSourceSonarConfig{
			ProjectKeys: ds.Spec.Config.ProjectKeys,
			Url:         ds.Spec.Config.Url,
			Username:    username,
			Password:    password,
		},
	}
}

func getSonarDsUpdateCommand(ds *v1alpha1.PerfDataSource, dsReq dto.DataSource) DataSourceCommand {
	return DataSourceCommand{
		Id:   dsReq.Id,
		Name: dsReq.Name,
		Type: DataSourceType(strings.ToUpper(dsReq.Type)),
		Config: DataSourceSonarConfig{
			ProjectKeys: append(dsReq.Config["projectKeys"].([]string), ds.Spec.Config.ProjectKeys...),
			Url:         dsReq.Config["url"].(string),
			Username:    dsReq.Config["username"].(string),
		},
	}
}

func getJenkinsDsCreateCommand(ds *v1alpha1.PerfDataSource, username string, password string) DataSourceCommand {
	return DataSourceCommand{
		Name: ds.Spec.Name,
		Type: DataSourceType(strings.ToUpper(ds.Spec.Type)),
		Config: DataSourceJenkinsConfig{
			JobNames: ds.Spec.Config.JobNames,
			Url:      ds.Spec.Config.Url,
			Username: username,
			Password: password,
		},
	}
}

func getJenkinsDsUpdateCommand(ds *v1alpha1.PerfDataSource, dsReq dto.DataSource) DataSourceCommand {
	return DataSourceCommand{
		Id:   dsReq.Id,
		Name: dsReq.Name,
		Type: DataSourceType(strings.ToUpper(dsReq.Type)),
		Config: DataSourceJenkinsConfig{
			JobNames: append(dsReq.Config["jobNames"].([]string), ds.Spec.Config.JobNames...),
			Url:      dsReq.Config["url"].(string),
			Username: dsReq.Config["username"].(string),
		},
	}
}
