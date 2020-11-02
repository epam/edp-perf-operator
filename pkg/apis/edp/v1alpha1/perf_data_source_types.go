package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerfDataSourceSpec defines the desired state of PerfDataSource
// +k8s:openapi-gen=true
type PerfDataSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Config         DataSourceConfig `json:"config"`
	PerfServerName string           `json:"perfServerName"`
}

type DataSourceConfig struct {
	JobNames       []string `json:"jobNames"`
	ProjectKeys    []string `json:"projectKeys"`
	Url            string   `json:"url"`
	CredentialName string   `json:"credentialName"`
}

// PerfDataSourceStatus defines the observed state of PerfDataSource
// +k8s:openapi-gen=true

type PerfDataSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSource is the Schema for the perfdatasources API
// +k8s:openapi-gen=true
type PerfDataSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceSpec   `json:"spec,omitempty"`
	Status PerfDataSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceList contains a list of PerfDataSource
type PerfDataSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerfDataSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSource{}, &PerfDataSourceList{})
}
