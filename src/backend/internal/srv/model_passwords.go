package srv

import "github.com/fedev521/g8keeper/backend/internal/types"

type CreatePasswordReqBody struct {
	Name   string `json:"name,omitempty"`
	Secret string `json:"password,omitempty"`
}

type ListPasswordsResponse200 struct {
	PasswordMetadata []types.PasswordMetadata `json:"passwords,omitempty"`
}

type GetPasswordResponse200 struct {
	Password types.Password `json:"password"`
}
