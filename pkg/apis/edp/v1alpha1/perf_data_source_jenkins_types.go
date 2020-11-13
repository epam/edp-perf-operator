package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerfDataSourceJenkinsSpec defines the desired state of PerfDataSource
// +k8s:openapi-gen=true
type PerfDataSourceJenkinsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
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

// PerfDataSourceJenkinsStatus defines the observed state of PerfDataSource
// +k8s:openapi-gen=true

type PerfDataSourceJenkinsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceJenkins is the Schema for the perfdatasourcejenkinses API
// +k8s:openapi-gen=true
type PerfDataSourceJenkins struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceJenkinsSpec   `json:"spec,omitempty"`
	Status PerfDataSourceJenkinsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceJenkinsList contains a list of PerfDataSource
type PerfDataSourceJenkinsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerfDataSourceJenkins `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceJenkins{}, &PerfDataSourceJenkinsList{})
}
