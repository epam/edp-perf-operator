package v1

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PerfServerSpec defines the desired state of PerfServer
type PerfServerSpec struct {
	ApiUrl         string `json:"apiUrl"`
	RootUrl        string `json:"rootUrl"`
	CredentialName string `json:"credentialName"`
	ProjectName    string `json:"projectName"`
}

type PerfServerStatus struct {
	// This flag indicates neither Codebase are initialized and ready to work. Defaults to false.
	Available bool `json:"available"`

	// Information when the last time the action were performed.
	LastTimeUpdated metaV1.Time `json:"last_time_updated"`

	// Detailed information regarding action result
	// which were performed
	// +optional
	DetailedMessage string `json:"detailed_message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ps
// +kubebuilder:storageversion

// PerfServer is the Schema for the PerfServers API
type PerfServer struct {
	metaV1.TypeMeta   `json:",inline"`
	metaV1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerfServerSpec   `json:"spec,omitempty"`
	Status PerfServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PerfServerList contains a list of PerfServer
type PerfServerList struct {
	metaV1.TypeMeta `json:",inline"`
	metaV1.ListMeta `json:"metadata,omitempty"`

	Items []PerfServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PerfServer{}, &PerfServerList{})
}
