package v1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource.
type PerfDataSourceJenkinsSpec struct {
	Name           string                  `json:"name"`
	Type           string                  `json:"type"`
	Config         DataSourceJenkinsConfig `json:"config"`
	PerfServerName string                  `json:"perfServerName"`
	CodebaseName   string                  `json:"codebaseName"`
}

type DataSourceJenkinsConfig struct {
	JobNames []string `json:"jobNames"`
	Url      string   `json:"url"`
}

// PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource.
type PerfDataSourceJenkinsStatus struct {
	// Specifies a current status of PerfDataSourceJenkins.
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pdsj,path=perfdatasourcejenkinses
// +kubebuilder:storageversion

// PerfDataSourceJenkins is the Schema for the PerfDataSourceJenkinses API.
type PerfDataSourceJenkins struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceJenkinsSpec   `json:"spec,omitempty"`
	Status PerfDataSourceJenkinsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerfDataSourceJenkinsList contains a list of PerfDataSource.
type PerfDataSourceJenkinsList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []PerfDataSourceJenkins `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceJenkins{}, &PerfDataSourceJenkinsList{})
}
