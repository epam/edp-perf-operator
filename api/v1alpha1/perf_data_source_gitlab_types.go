package v1alpha1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab.
type PerfDataSourceGitLabSpec struct {
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Config         DataSourceGitLabConfig `json:"config"`
	PerfServerName string                 `json:"perfServerName"`
	CodebaseName   string                 `json:"codebaseName"`
}

type DataSourceGitLabConfig struct {
	Repositories []string `json:"repositories"`
	Url          string   `json:"url"`
	Branches     []string `json:"branches"`
}

type PerfDataSourceGitLabStatus struct {
	// Specifies a current status of PerfDataSourceGitLab.
	Status string `json:"status"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=pdsgl
// +kubebuilder:deprecatedversion

// PerfDataSourceGitLab is the Schema for the PerfDataSourceGitLabs API.
type PerfDataSourceGitLab struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceGitLabSpec   `json:"spec,omitempty"`
	Status PerfDataSourceGitLabStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerfDataSourceGitLabList contains a list of PerfDataSourceGitLab.
type PerfDataSourceGitLabList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []PerfDataSourceGitLab `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceGitLab{}, &PerfDataSourceGitLabList{})
}
