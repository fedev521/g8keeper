package crypt

import (
	"fmt"

	"github.com/fedev521/tinksrv/internal/log"
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/testing/fakekms"
)

// The fake KMS should only be used in tests. It is not secure.
const keyURI = "fake-kms://CM2b3_MDElQKSAowdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUuY3J5cHRvLnRpbmsuQWVzR2NtS2V5EhIaEIK75t5L-adlUwVhWvRuWUwYARABGM2b3_MDIAE"

// When you encrypt data with the KmsEnvelopeAead primitive, Tink:
// 1. Generates a random DEK and locally encrypts your data with it.
// 2. Makes a request to your KMS to encrypt the DEK with the KMS KEK.
// 3. Concatenates the KEK-encrypted encryption DEK with the encrypted data.
// Both the DEK type and the KEK URI are specified in the KmsEnvelopeAead key.

func EnvelopeEncrypt(plaintext, associatedData []byte) ([]byte, error) {
	// Get a KEK (key encryption key) AEAD. This is usually a remote AEAD to a
	// KMS. In this example, we use a fake KMS to avoid making RPCs.
	client, err := fakekms.NewClient(keyURI)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("could not create KMS client: %w", err)
	}
	kekAEAD, err := client.GetAEAD(keyURI)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("could not get AEAD backend: %w", err)
	}

	// Get the KMS envelope AEAD primitive.
	primitive := aead.NewKMSEnvelopeAEAD2(aead.AES256GCMKeyTemplate(), kekAEAD)

	ciphertext, err := primitive.Encrypt(plaintext, associatedData)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("error during encryption: %w", err)
	}

	return ciphertext, nil
}

// When decrypting, Tink does the reverse operations:
// 1. Extracts the KEK-encrypted DEK key.
// 2. Makes a request to your KMS to decrypt the KEK-encrypted DEK.
// 3. Decrypts the ciphertext locally using the DEK.

func EnvelopeDecrypt(ciphertext, associatedData []byte) ([]byte, error) {
	// Get a KEK (key encryption key) AEAD. This is usually a remote AEAD to a
	// KMS. In this example, we use a fake KMS to avoid making RPCs.
	client, err := fakekms.NewClient(keyURI)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("could not create KMS client: %w", err)
	}
	kekAEAD, err := client.GetAEAD(keyURI)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("could not get AEAD backend: %w", err)
	}

	// Get the KMS envelope AEAD primitive.
	primitive := aead.NewKMSEnvelopeAEAD2(aead.AES256GCMKeyTemplate(), kekAEAD)

	plaintext, err := primitive.Decrypt(ciphertext, associatedData)
	if err != nil {
		log.Error(err.Error())
		return []byte{}, fmt.Errorf("error during decryption: %w", err)
	}

	return plaintext, nil
}
