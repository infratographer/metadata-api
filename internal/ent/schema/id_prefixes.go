package schema

const (
	// ApplicationPrefix is the prefix for all application IDs owned by metadata-api
	ApplicationPrefix string = "meta"
	// MetadataPrefix is the prefix for metadata
	MetadataPrefix string = ApplicationPrefix + "dat"
	// AnnotationPrefix is the prefix for all annotations
	AnnotationPrefix string = ApplicationPrefix + "ano"
	// StatusPrefix is the prefix for all statuses
	StatusPrefix string = ApplicationPrefix + "sts"
	// MetadataNamespacePrefix is the prefix for all metadata namespaces
	MetadataNamespacePrefix string = ApplicationPrefix + "mns"
	// StatusNamespacePrefix is the prefix for all status namespaces
	StatusNamespacePrefix string = ApplicationPrefix + "sns"
)
