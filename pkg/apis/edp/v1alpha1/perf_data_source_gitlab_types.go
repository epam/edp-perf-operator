package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerfDataSourceGitLabSpec defines the desired state of PerfDataSourceGitLab
// +k8s:openapi-gen=true
type PerfDataSourceGitLabSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
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

// PerfDataSourceGitLabStatus defines the observed state of PerfDataSourceGitLab
// +k8s:openapi-gen=true

type PerfDataSourceGitLabStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceGitLab is the Schema for the perfdatasourcegitlabs API
// +k8s:openapi-gen=true
type PerfDataSourceGitLab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceGitLabSpec   `json:"spec,omitempty"`
	Status PerfDataSourceGitLabStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceGitLabList contains a list of PerfDataSourceGitLab
type PerfDataSourceGitLabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerfDataSourceGitLab `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceGitLab{}, &PerfDataSourceGitLabList{})
}
