package spec

// ManagedClusterMetadata struct holds information for managed cluster metadata.
type ManagedClusterMetadata struct {
	Name             string   `json:"name"`
	Namespace        string   `json:"namespace"`
	Labels           []string `json:"labels"`
	DeletedLabelKeys []string `json:"deletedLabelKeys"`
	Version          uint64   `json:"version"`
}
