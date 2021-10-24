package datatypes

const (
	// SpecBundle - spec bundle message type.
	SpecBundle = "SpecBundle"
	// StatusBundle - status bundle message type.
	StatusBundle = "StatusBundle"
	// Config - hub of hubs config message type.
	Config = "Config"

	// ManagedClustersMsgKey - managed clusters message key.
	ManagedClustersMsgKey = "ManagedClusters"
	// ClustersPerPolicyMsgKey - clusters per policy message key.
	ClustersPerPolicyMsgKey = "ClustersPerPolicy"
	// PolicyCompleteComplianceMsgKey - policy complete state compliance message key.
	PolicyCompleteComplianceMsgKey = "PolicyCompleteCompliance"
	// PolicyDeltaComplianceMsgKey - policy delta state compliance message key.
	PolicyDeltaComplianceMsgKey = "PolicyDeltaCompliance"
	// MinimalPolicyComplianceMsgKey - minimal policy compliance message key.
	MinimalPolicyComplianceMsgKey = "MinimalPolicyCompliance"

	// TODO to be deleted after merging kafka changes, for backward compatibility until other repos adapt new types

	// PolicyComplianceMsgKey - policy compliance message key.
	PolicyComplianceMsgKey = "PolicyCompliance"
	// LocalPlacementRulesMsgKey - local placement rules message key.
	LocalPlacementRulesMsgKey = "LocalPlacementRules"
	// LocalClustersPerPolicyMsgKey - local clusters per policy message key.
	LocalClustersPerPolicyMsgKey = "LocalClustersPerPolicy"
	// LocalPolicyComplianceMsgKey - local policy compliance message key.
	LocalPolicyComplianceMsgKey = "LocalPolicyCompliance"
	// LocalSpecPerPolicyMsgKey - the local spec per policy message key.
	LocalSpecPerPolicyMsgKey = "localSpecPerPolicy"
)
