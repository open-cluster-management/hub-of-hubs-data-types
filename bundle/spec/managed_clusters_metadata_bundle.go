package spec

// ManagedClustersMetadata struct bundles ManagedClusterMetadata objects.
type ManagedClustersMetadata struct {
	Objects     []*ManagedClusterMetadata `json:"objects"`
	LeafHubName string                    `json:"leafHubName"`
}
