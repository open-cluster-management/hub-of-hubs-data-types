package status

// PolicyGenericComplianceStatus struct holds information for policy compliance status and may be used either for
// complete or delta state bundles.
type PolicyGenericComplianceStatus struct {
	PolicyID                  string   `json:"policyId"`
	CompliantClusters         []string `json:"compliantClusters"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
}

// PolicyCompleteComplianceStatus struct holds information for (optimized) policy compliance status.
type PolicyCompleteComplianceStatus struct {
	PolicyID                  string   `json:"policyId"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
}

// BaseClustersPerPolicyBundle is the base struct for clusters per policy bundle and contains the full state.
type BaseClustersPerPolicyBundle struct {
	Objects     []*PolicyGenericComplianceStatus `json:"objects"`
	LeafHubName string                           `json:"leafHubName"`
	BundleVersion
}

// BaseCompleteComplianceStatusBundle is the base struct for complete state compliance status bundle.
type BaseCompleteComplianceStatusBundle struct {
	Objects              []*PolicyCompleteComplianceStatus `json:"objects"`
	LeafHubName          string                            `json:"leafHubName"`
	BaseBundleGeneration uint64                            `json:"baseBundleGeneration"`
	BundleVersion
}

// BaseDeltaComplianceStatusBundle is the base struct for delta state compliance status bundle.
type BaseDeltaComplianceStatusBundle struct {
	Objects              []*PolicyGenericComplianceStatus `json:"objects"`
	LeafHubName          string                           `json:"leafHubName"`
	BaseBundleGeneration uint64                           `json:"baseBundleGeneration"`
	BundleVersion
}
