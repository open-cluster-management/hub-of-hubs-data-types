package datatypes

const (
	// HohSystemNamespace - Hub of Hubs dedicated namespace.
	HohSystemNamespace = "hoh-system"

	// OriginOwnerReferenceAnnotation - Origin owner reference annotation.
	OriginOwnerReferenceAnnotation = "hub-of-hubs.open-cluster-management.io/originOwnerReferenceUid"

	// HohSecretAnno - mark the secret as a Hoh related resource, meaning hoh controllers will act on secrets
	// with this annotation
	HohSecretAnnotation = "hub-of-hubs.open-cluster-management.io/secret"
)
