package store

import "github.com/fedev521/g8keeper/backend/internal/types"

type PasswordKeeper interface {
	Store(p types.Password, key string) error
	Retrieve(key string) (types.Password, error)
	ListMetadata() ([]types.PasswordMetadata, error)
}
