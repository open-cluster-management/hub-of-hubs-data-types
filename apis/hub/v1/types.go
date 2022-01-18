// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HubSpec defines the desired state of Hub.
type HubSpec struct {
	UUID string `json:"uuid"`
}

// HubStatus defines the observed state of Hub.
type HubStatus struct {
	// Conditions represent the latest available observations of an object's state
	Conditions                []metav1.Condition `json:"conditions,omitempty"`
	TransportBootstrapServers string             `json:"transport-bootstrap-servers,omitempty"`
	TransportCertificate      string             `json:"transport-certificate,omitempty"`
	TransportSpecTopic        string             `json:"transport-spec-topic,omitempty"`
	WriterIncarnation         string             `json:"writer-incarnation,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hub is the Schema for the hubs API.
type Hub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HubSpec   `json:"spec,omitempty"`
	Status HubStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HubList contains a list of Config
type HubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hub `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hub{}, &HubList{})
}
