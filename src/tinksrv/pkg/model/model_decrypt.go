package model

type DecryptReqBody struct {
	// The ciphertext to decrypt.
	CiphertextB64 string `json:"ciphertext,omitempty"`
}

type DecryptResponse200 struct {
	// The decrypted data.
	Plaintext string `json:"plaintext,omitempty"`
}
