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
	// ManagedClustersMetadataMsgKey - managed clusters metadata message key.
	ManagedClustersMetadataMsgKey = "ManagedClustersMetadata"
	// ManagedClusterSetsMsgKey - managed cluster sets message key.
	ManagedClusterSetsMsgKey = "ManagedClusterSets"
	// ManagedClusterSetBindingsMsgKey - managed cluster set bindings message key.
	ManagedClusterSetBindingsMsgKey = "ManagedClusterSetBindings"

	// ClustersPerPolicyMsgKey - clusters per policy message key.
	ClustersPerPolicyMsgKey = "ClustersPerPolicy"
	// PolicyCompleteComplianceMsgKey - policy complete state compliance message key.
	PolicyCompleteComplianceMsgKey = "PolicyCompleteCompliance"
	// PolicyDeltaComplianceMsgKey - policy delta state compliance message key.
	PolicyDeltaComplianceMsgKey = "PolicyDeltaCompliance"
	// PolicyPlacementMsgKey - policy placement message key.
	PolicyPlacementMsgKey = "PolicyPlacement"
	// MinimalPolicyComplianceMsgKey - minimal policy compliance message key.
	MinimalPolicyComplianceMsgKey = "MinimalPolicyCompliance"
	// PoliciesMsgKey - the policy spec message key.
	PoliciesMsgKey = "Policies"
	// PlacementsMsgKey - placements message key.
	PlacementsMsgKey = "Placements"
	// PlacementRulesMsgKey - placement rules message key.
	PlacementRulesMsgKey = "PlacementRules"
	// PlacementBindingsMsgKey - placement bindings message key.
	PlacementBindingsMsgKey = "PlacementBindings"

	// LocalPolicySpecMsgKey - the local policy spec message key.
	LocalPolicySpecMsgKey = "LocalPolicySpec"
	// LocalPlacementRulesMsgKey - local placement rules message key.
	LocalPlacementRulesMsgKey = "LocalPlacementRules"
	// LocalClustersPerPolicyMsgKey - local clusters per policy message key.
	LocalClustersPerPolicyMsgKey = "LocalClustersPerPolicy"
	// LocalPolicyCompleteComplianceMsgKey - local policy compliance message key.
	LocalPolicyCompleteComplianceMsgKey = "LocalPolicyCompleteCompliance"

	// ChannelsMsgKey - the channels spec message key.
	ChannelsMsgKey = "Channels"
	// ApplicationsMsgKey - the applications spec message key.
	ApplicationsMsgKey = "Applications"
	// SubscriptionsMsgKey - the subscriptions spec message key.
	SubscriptionsMsgKey = "Subscriptions"
	// SubscriptionStatusMsgKey - subscription status message key.
	SubscriptionStatusMsgKey = "SubscriptionStatus"

	// ControlInfoMsgKey - control info message key.
	ControlInfoMsgKey = "ControlInfo"
)
