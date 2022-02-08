package spec

// ManagedClustersMetadata struct bundles ManagedClusterMetadata objects.
type ManagedClustersMetadata struct {
	Objects     []*ManagedClusterMetadata `json:"objects"`
	LeafHubName string                    `json:"leafHubName"`
	Version     uint64                    `json:"version"` // highest inner object version
}
