package status

// BundleVersion holds the information necessary for the consumers of status bundles to compare versions correctly.
// Incarnation is the instance-version of the sender component, a counter that increments on restarts of the latter.
// Generation is the bundle-version relative to the sender's instance, increments on bundle updates.
type BundleVersion struct {
	Incarnation uint64 `json:"incarnation"`
	Generation  uint64 `json:"generation"`
}

// NewerThan returns whether the caller's version is newer than that received as argument.
func (bv *BundleVersion) NewerThan(other *BundleVersion) bool {
	if bv.Incarnation == other.Incarnation {
		return bv.Generation > other.Generation
	}

	return bv.Incarnation > other.Incarnation
}
