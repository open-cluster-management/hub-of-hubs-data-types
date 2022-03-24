package status

// PolicyPlacementStatus holds information for policy placement status.
type PolicyPlacementStatus struct {
	PolicyID         string `json:"policyId"`
	PlacementRule    string `json:"placementRule"`
	PlacementBinding string `json:"placementBinding"`
}

// BasePolicyPlacementStatusBundle the base struct for minimal compliance status bundle.
type BasePolicyPlacementStatusBundle struct {
	Objects       []*PolicyPlacementStatus `json:"objects"`
	LeafHubName   string                   `json:"leafHubName"`
	BundleVersion *BundleVersion           `json:"bundleVersion"`
}
