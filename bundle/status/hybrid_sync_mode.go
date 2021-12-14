package status

// HybridSyncMode used to identify hybrid sync mode - complete/delta bundles.
type HybridSyncMode int8

const (
	// CompleteStateMode used to identify sync mode of complete state bundles.
	CompleteStateMode HybridSyncMode = iota
	// DeltaStateMode used to identify sync mode of delta state bundles.
	DeltaStateMode HybridSyncMode = iota
)
