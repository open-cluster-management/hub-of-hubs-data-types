// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AggregationLevel specifies the level of aggregation leaf hubs should do before sending the information
type AggregationLevel string

const (
	// Full is an AggregationLevel
	Full AggregationLevel = "full"

	// Minimal is an AggregationLevel
	Minimal AggregationLevel = "minimal"

	// ShowLocalPolicies decides if to propagate local policies
	ShowLocalPolicies bool = true
)

// ConfigSpec defines the desired state of Config
type ConfigSpec struct {
	AggregationLevel AggregationLevel `json:"aggregationLevel,omitempty"` // full or minimal
}

// ConfigStatus defines the observed state of Config
type ConfigStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Config is the Schema for the configs API
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigList contains a list of Config
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Config `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Config{}, &ConfigList{})
}
