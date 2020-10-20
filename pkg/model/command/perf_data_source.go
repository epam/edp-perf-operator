package command

import "github.com/epmd-edp/perf-operator/pkg/apis/edp/v1alpha1"

type DeactivationReason string

const (
	AutoDueToErrors DeactivationReason = "AUTO_DUE_TO_ERRORS"
	Manual          DeactivationReason = "MANUAL"
)

type DataSourceType string

const (
	Jira  DataSourceType = "JIRA"
	Sonar DataSourceType = "SONAR"
)

type DataSourceCreateCommand struct {
	Active                bool               `json:"active"`
	Config                DataSourceConfig   `json:"config"`
	DeactivationReason    DeactivationReason `json:"deactivationReason"`
	DeactivationTimestamp string             `json:"deactivationTimestamp"`
	DeprecationDate       string             `json:"deprecationDate"`
	Id                    int                `json:"id"`
	MetaData              DataSourceMetaData `json:"metaData"`
	Name                  string             `json:"name"`
	Type                  DataSourceType     `json:"type"`
	WorkflowUpdateRequire bool               `json:"workflowUpdateRequire"`
}

type DataSourceConfig struct {
	InstanceId string `json:"instanceId"`
}

type DataSourceMetaData struct {
}

func ConvertToDataSourceCreateCommandModel(ds *v1alpha1.PerfDataSource) DataSourceCreateCommand {
	return DataSourceCreateCommand{
		Active:                true,
		Config:                DataSourceConfig{},
		DeactivationReason:    Manual,
		DeactivationTimestamp: "",
		DeprecationDate:       "",
		Id:                    0,
		MetaData:              DataSourceMetaData{},
		Name:                  ds.Name,
		Type:                  DataSourceType(ds.Spec.Type),
		WorkflowUpdateRequire: true,
	}
}
