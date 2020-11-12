package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerfDataSourceSonarSpec defines the desired state of PerfDataSourceSonar
// +k8s:openapi-gen=true
type PerfDataSourceSonarSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	Config         DataSourceSonarConfig `json:"config"`
	PerfServerName string                `json:"perfServerName"`
}

type DataSourceSonarConfig struct {
	ProjectKeys []string `json:"projectKeys"`
	Url         string   `json:"url"`
}

// PerfDataSourceSonartatus defines the observed state of PerfDataSourceSonar
// +k8s:openapi-gen=true

type PerfDataSourceSonarStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceSonar is the Schema for the perfdatasourcessonars API
// +k8s:openapi-gen=true
type PerfDataSourceSonar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfDataSourceSonarSpec   `json:"spec,omitempty"`
	Status PerfDataSourceSonarStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerfDataSourceSonarList contains a list of PerfDataSourceSonar
type PerfDataSourceSonarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerfDataSourceSonar `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfDataSourceSonar{}, &PerfDataSourceSonarList{})
}
