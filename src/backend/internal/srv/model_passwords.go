package srv

import "github.com/fedev521/g8keeper/backend/internal/types"

type ListPasswordsResponse200 struct {
	PasswordMetadata []types.PasswordMetadata `json:"passwords,omitempty"`
}

type CreatePasswordReqBody struct {
	Name   string `json:"name,omitempty"`
	Secret string `json:"password,omitempty"`
}
