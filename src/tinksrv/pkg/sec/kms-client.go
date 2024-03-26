package sec

import (
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"github.com/tink-crypto/tink-go/v2/tink"
)

// G8KeeperKMSClient acts as a KMSClient, therefore it knows how to produce
// primitives backed by keys stored in remote KMS services.
type G8KeeperKMSClient struct{}

func NewG8KeeperKMSClient() *G8KeeperKMSClient {
	return &G8KeeperKMSClient{}
}

func (k G8KeeperKMSClient) Supported(keyURI string) bool {
	// TODO
	return true
}

func (k G8KeeperKMSClient) GetAEAD(keyURI string) (tink.AEAD, error) {
	// TODO
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		return &AlwaysFailingAead{Error: err}, err
	}
	a, err := aead.New(kh)
	if err != nil {
		return &AlwaysFailingAead{Error: err}, err
	}
	return a, nil
}
