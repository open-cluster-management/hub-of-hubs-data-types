// Copyright (c) 2021 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HohHeartbeatStatus defines the observed state of HohHeartbeat
type HohHeartbeatStatus struct {
	// LastTimestamp contains last timestamp a heartbeat was received from HoH.
	// Format: number of seconds elapsed since January 1, 1970 UTC.
	LastTimestamp int64 `json:"lastTimestamp,omitempty"`
	// Conditions contains the different condition statuses for this HoH heartbeat.
	Conditions []metav1.Condition `json:"conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HohHeartbeat is the Schema for the hoh heartbeats API
type HohHeartbeat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status HohHeartbeatStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HohHeartbeatList contains a list of HohHeartbeat
type HohHeartbeatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HohHeartbeat `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HohHeartbeat{}, &HohHeartbeatList{})
}
