package spec

import "time"

// ManagedClusterLabels struct holds information for managed cluster metadata.
type ManagedClusterLabels struct {
	Name             string            `json:"name"`
	Namespace        string            `json:"namespace"`
	Labels           map[string]string `json:"labels"`
	DeletedLabelKeys []string          `json:"deletedLabelKeys"`
	UpdateTimestamp  time.Time         `json:"updateTimestamp"`
}
