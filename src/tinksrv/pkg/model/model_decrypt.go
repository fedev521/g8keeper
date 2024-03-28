package model

type DecryptReqBody struct {
	// The ciphertext to decrypt.
	CiphertextB64 string `json:"ciphertext,omitempty"`
	// TODO add aad
}

type DecryptResponse200 struct {
	// The decrypted data.
	Plaintext string `json:"plaintext,omitempty"`
}
