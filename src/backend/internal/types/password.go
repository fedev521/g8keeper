package types

type PasswordMetadata struct {
	Name  string   `json:"name,omitempty"`
	User  string   `json:"user,omitempty"`
	Sites []string `json:"sites,omitempty"`
}

type Password struct {
	Secret   string           `json:"secret,omitempty"`
	Metadata PasswordMetadata `json:"metadata,omitempty"`
}
