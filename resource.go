package coinbase

type Resource struct {
	// Resource ID
	ID string `json:"id,omitempty"`
	// Resource name
	Resource string `json:"resource,omitempty"`
	// Resource path
	ResourcePath string `json:"resource_path,omitempty"`
}
