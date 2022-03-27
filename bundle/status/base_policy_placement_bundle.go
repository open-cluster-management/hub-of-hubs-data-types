package status

import policiesv1 "github.com/open-cluster-management/governance-policy-propagator/api/v1"

// PolicyPlacementStatus holds information for policy placement status.
type PolicyPlacementStatus struct {
	PolicyID  string                  `json:"policyId"`
	Placement []*policiesv1.Placement `json:"placement"`
}

// BasePolicyPlacementStatusBundle the base struct for policy placement status bundle.
type BasePolicyPlacementStatusBundle struct {
	Objects       []*PolicyPlacementStatus `json:"objects"`
	LeafHubName   string                   `json:"leafHubName"`
	BundleVersion *BundleVersion           `json:"bundleVersion"`
}
