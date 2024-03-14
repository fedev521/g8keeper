package srv

type PasswordMetadata struct {
	Name string `json:"name,omitempty"`
}

type ListPasswordsResponse200 struct {
	PasswordMetadata []PasswordMetadata `json:"passwords,omitempty"`
}
