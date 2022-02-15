package spec

import "time"

// ManagedClusterLabelsSpec struct holds information for managed cluster metadata.
type ManagedClusterLabelsSpec struct {
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	Labels           map[string]string `json:"labels"`
	DeletedLabelKeys []string          `json:"deletedLabelKeys"`
	UpdateTimestamp  time.Time         `json:"updateTimestamp"`
}
