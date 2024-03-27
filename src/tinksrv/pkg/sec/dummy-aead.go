package sec

import "fmt"

type AlwaysFailingAead struct {
	Error error
}

// Encrypt returns an error on encryption.
func (a *AlwaysFailingAead) Encrypt(plaintext []byte, associatedData []byte) ([]byte, error) {
	return nil, fmt.Errorf("AlwaysFailingAead will always fail on encryption: %w", a.Error)
}

// Decrypt returns an error on decryption.
func (a *AlwaysFailingAead) Decrypt(ciphertext []byte, associatedData []byte) ([]byte, error) {
	return nil, fmt.Errorf("AlwaysFailingAead will always fail on decryption: %w", a.Error)
}
