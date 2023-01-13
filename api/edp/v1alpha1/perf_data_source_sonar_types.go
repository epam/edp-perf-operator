package v1alpha1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar.
type PerfDataSourceSonarSpec struct {
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	Config         DataSourceSonarConfig `json:"config"`
	PerfServerName string                `json:"perfServerName"`
	CodebaseName   string                `json:"codebaseName"`
}

type DataSourceSonarConfig struct {
	ProjectKeys []string `json:"projectKeys"`
	Url         string   `json:"url"`
}

// PerfDataSourceSonarStatus defines the observed state of PerfDataSourceSonar.
type PerfDataSourceSonarStatus struct {
	// Specifies a current status of PerfDataSourceSonar.
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pdss
// +kubebuilder:deprecatedversion

// PerfDataSourceSonar is the Schema for the PerfDataSourceSonars API.
type PerfDataSourceSonar struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceSonarSpec   `json:"spec,omitempty"`
	Status PerfDataSourceSonarStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerfDataSourceSonarList contains a list of PerfDataSourceSonar.
type PerfDataSourceSonarList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []PerfDataSourceSonar `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceSonar{}, &PerfDataSourceSonarList{})
}
