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
	// HeartbeatKey - heartbeat message key.
	HeartbeatKey = "Heartbeat"
)
