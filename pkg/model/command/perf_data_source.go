package command

import (
	"github.com/epmd-edp/perf-operator/v2/pkg/apis/edp/v1alpha1"
	"github.com/epmd-edp/perf-operator/v2/pkg/model/dto"
	"github.com/epmd-edp/perf-operator/v2/pkg/util/common"
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

type DataSourceConfigDto struct {
	Type       string
	ApiUrl     string
	Username   string
	Password   string
	Parameters []string
}

func GetSonarDsCreateCommand(ds *v1alpha1.PerfDataSourceSonar, username string, password string) DataSourceCommand {
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

func GetSonarDsUpdateCommand(dsReq *dto.DataSource, conf DataSourceConfigDto) DataSourceCommand {
	return DataSourceCommand{
		Id:   dsReq.Id,
		Name: dsReq.Name,
		Type: DataSourceType(strings.ToUpper(dsReq.Type)),
		Config: DataSourceSonarConfig{
			ProjectKeys: append(common.ConvertToStringArray(dsReq.Config["projectKeys"]), conf.Parameters...),
			Url:         conf.ApiUrl,
			Username:    conf.Username,
			Password:    conf.Password,
		},
	}
}

func GetJenkinsDsCreateCommand(ds *v1alpha1.PerfDataSourceJenkins, username string, password string) DataSourceCommand {
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

func GetJenkinsDsUpdateCommand(dsReq *dto.DataSource, conf DataSourceConfigDto) DataSourceCommand {
	return DataSourceCommand{
		Id:   dsReq.Id,
		Name: dsReq.Name,
		Type: DataSourceType(strings.ToUpper(dsReq.Type)),
		Config: DataSourceJenkinsConfig{
			JobNames: append(common.ConvertToStringArray(dsReq.Config["jobNames"]), conf.Parameters...),
			Url:      conf.ApiUrl,
			Username: conf.Username,
			Password: conf.Password,
		},
	}
}
