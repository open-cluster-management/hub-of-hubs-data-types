package status

// PolicyCompleteComplianceStatus struct holds information for (optimized) policy compliance status.
type PolicyCompleteComplianceStatus struct {
	PolicyID                  string   `json:"policyId"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
	ResourceVersion           string   `json:"resourceVersion"`
}

// PolicyDeltaComplianceStatus struct holds information for raw policy compliance status.
type PolicyDeltaComplianceStatus struct {
	PolicyID                  string   `json:"policyId"`
	CompliantClusters         []string `json:"compliantClusters"`
	NonCompliantClusters      []string `json:"nonCompliantClusters"`
	UnknownComplianceClusters []string `json:"unknownComplianceClusters"`
	ResourceVersion           string   `json:"resourceVersion"`
}

// BaseCompleteComplianceStatusBundle is the base struct for complete state compliance status bundle.
type BaseCompleteComplianceStatusBundle struct {
	Objects              []*PolicyCompleteComplianceStatus `json:"objects"`
	LeafHubName          string                            `json:"leafHubName"`
	BaseBundleGeneration uint64                            `json:"baseBundleGeneration"`
	Generation           string                            `json:"generation"`
}

// BaseDeltaComplianceStatusBundle is the base struct for delta state compliance status bundle.
type BaseDeltaComplianceStatusBundle struct {
	Objects              []*PolicyDeltaComplianceStatus `json:"objects"`
	LeafHubName          string                         `json:"leafHubName"`
	BaseBundleGeneration uint64                         `json:"baseBundleGeneration"`
	Generation           string                         `json:"generation"`
}
