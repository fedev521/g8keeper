package store

import (
	"errors"

	"github.com/fedev521/g8keeper/backend/internal/types"
)

type InMapKeeper struct {
	memory map[string]types.Password
}

func NewInMapKeeper() *InMapKeeper {
	return &InMapKeeper{
		memory: make(map[string]types.Password),
	}
}

func (k InMapKeeper) Store(p types.Password, key string) error {
	k.memory[key] = p
	return nil
}

func (k InMapKeeper) Retrieve(key string) (types.Password, error) {
	p, found := k.memory[key]
	if found {
		return p, nil
	}
	return types.Password{}, errors.New("no password matching key")
}

func (k InMapKeeper) ListMetadata() ([]types.PasswordMetadata, error) {
	metadata := make([]types.PasswordMetadata, 0, len(k.memory))
	for _, p := range k.memory {
		metadata = append(metadata, p.Metadata)
	}
	return metadata, nil
}
