package meta

// ServiceMeta contains URLs for metadata regarding the service
type ServiceMeta struct {

	// VersionPrefix defines the prefix used to prefix routes for this specific version
	VersionPrefix string `json:"version_prefix,omitempty"`

	// URL to the API specification
	Specification string `json:"specification,omitempty"`
}
