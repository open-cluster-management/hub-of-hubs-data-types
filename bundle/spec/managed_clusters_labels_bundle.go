package spec

// ManagedClusterLabelsSpecBundle struct bundles ManagedClusterLabelsSpec objects.
type ManagedClusterLabelsSpecBundle struct {
	Objects     []*ManagedClusterLabelsSpec `json:"objects"`
	LeafHubName string                      `json:"leafHubName"`
}
