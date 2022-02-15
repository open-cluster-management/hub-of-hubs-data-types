package spec

// ManagedClustersLabelsBundle struct bundles ManagedClusterLabels objects.
type ManagedClustersLabelsBundle struct {
	Objects     []*ManagedClusterLabels `json:"objects"`
	LeafHubName string                  `json:"leafHubName"`
}
