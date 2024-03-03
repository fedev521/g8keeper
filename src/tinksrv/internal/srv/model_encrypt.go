package srv

type EncryptReqBody struct {
	// The plaintext to encrypt.
	Plaintext string `json:"plaintext,omitempty"`
}

type EncryptResponse200 struct {
	// The encrypted data.
	CiphertextB64 string `json:"ciphertext,omitempty"`
}
