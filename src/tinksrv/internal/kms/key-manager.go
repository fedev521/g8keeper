package kms

import (
	"bufio"
	"os"

	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/insecurecleartextkeyset"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"github.com/tink-crypto/tink-go/v2/tink"
)

type KEKManager struct {
	primitive tink.AEAD
}

func NewKEKManager(conf Config) (*KEKManager, error) {
	kekFile, err := os.Open(conf.KekFile)
	if err != nil {
		return &KEKManager{}, err
	}
	defer kekFile.Close()
	kekReader := bufio.NewReader(kekFile)

	kh, err := insecurecleartextkeyset.Read(keyset.NewJSONReader(kekReader))
	if err != nil {
		return &KEKManager{}, err
	}
	primitive, err := aead.New(kh)
	if err != nil {
		return &KEKManager{}, err
	}

	return &KEKManager{primitive: primitive}, nil
}

func (km KEKManager) GetAEAD() (tink.AEAD, error) {
	return km.primitive, nil
}
