package store

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sync"

	"github.com/fedev521/g8keeper/backend/internal/log"
	"github.com/fedev521/g8keeper/backend/internal/svc"
	"github.com/fedev521/g8keeper/backend/internal/types"
	"github.com/fedev521/g8keeper/tinksrv/pkg/sec"
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/tink"
)

type CryptedInMemKeeper struct {
	memory    *sync.Map
	primitive tink.AEAD
}

func NewCryptedInMemKeeper(tinkSvcConf svc.TinkSvcConfig) (*CryptedInMemKeeper, error) {
	rka, err := sec.NewRemoteKEKAEAD(tinkSvcConf.Host, tinkSvcConf.Port)
	if err != nil {
		return nil, fmt.Errorf("could create remote KEK AEAD: %w", err)
	}
	primitive := aead.NewKMSEnvelopeAEAD2(aead.AES256GCMKeyTemplate(), rka)

	keeper := &CryptedInMemKeeper{
		memory:    new(sync.Map),
		primitive: primitive,
	}
	return keeper, nil
}

func (k CryptedInMemKeeper) Store(p types.Password, key string) error {
	pt := []byte(p.Secret)
	aad := []byte("aad")

	ct, err := k.primitive.Encrypt(pt, aad)
	if err != nil {
		return fmt.Errorf("error during encryption: %w", err)
	}
	// update in-place the secret with its encrypted base64-encoded form
	p.Secret = base64.StdEncoding.EncodeToString(ct)

	k.memory.Store(key, p)
	log.Debug("stored password", map[string]interface{}{
		"key": key,
	})
	return nil
}

func (k CryptedInMemKeeper) Retrieve(key string) (types.Password, error) {
	p, found := k.memory.Load(key)
	if !found {
		return types.Password{}, errors.New("no password matching key")
	}

	typedPassword, cast := p.(types.Password)
	if !cast {
		return types.Password{}, errors.New("cannot cast to Password type")
	}

	ct, err := base64.StdEncoding.DecodeString(typedPassword.Secret)
	if err != nil {
		return types.Password{}, fmt.Errorf("error during base64 decoding: %w", err)
	}
	aad := []byte("aad")
	pt, err := k.primitive.Decrypt(ct, aad)
	if err != nil {
		return types.Password{}, fmt.Errorf("error during decryption: %w", err)
	}

	typedPassword.Secret = string(pt)
	return typedPassword, nil
}

func (k CryptedInMemKeeper) ListMetadata() ([]types.PasswordMetadata, error) {
	metadata := make([]types.PasswordMetadata, 0)
	k.memory.Range(func(key, p any) bool {
		typedPassword, cast := p.(types.Password)
		if !cast {
			log.Error("unable to cast password to types.Password")
			return true
		}
		metadata = append(metadata, typedPassword.Metadata)
		return true
	})

	return metadata, nil
}
