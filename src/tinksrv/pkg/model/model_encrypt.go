package model

type EncryptReqBody struct {
	// The base64-encoded plaintext to encrypt.
	PlaintextB64 string `json:"plaintext,omitempty"`
	// TODO add aad
}

type EncryptResponse200 struct {
	// The base64-encoded encrypted data.
	CiphertextB64 string `json:"ciphertext,omitempty"`
}
