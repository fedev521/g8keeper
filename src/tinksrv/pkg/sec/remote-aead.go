package sec

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fedev521/g8keeper/tinksrv/pkg/model"
)

// RemoteKEKAEAD implements the tink.AEAD interface. It is meant to be used with
// envelope encryption, e.g. `aead.NewKMSEnvelopeAEAD2(kt, rka)`. It
// encrypts/decrypts data (note: a DEK) by leveraging an external service so
// that the KEK never leaves the KMS. Its Encrypt()/Decrypts methods are invoked
// by aead.NewKMSEnvelopeAEAD2() on the DEK.
type RemoteKEKAEAD struct {
	baseURL *url.URL
	Error   error
}

func NewRemoteKEKAEAD(host, port string) (*RemoteKEKAEAD, error) {
	baseURL, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	rka := &RemoteKEKAEAD{
		baseURL: baseURL,
	}
	return rka, nil
}

// Encrypt implements the tink.AEAD interface for encryption.
func (rka *RemoteKEKAEAD) Encrypt(dek []byte, _ []byte) ([]byte, error) {

	// prepare and send request
	reqURL := rka.baseURL.JoinPath("/v1/encrypt").String()
	reqBody, _ := json.Marshal(model.EncryptReqBody{
		PlaintextB64: base64.StdEncoding.EncodeToString(dek),
	})
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request to KMS: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send HTTP request to KMS: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not wrap DEK within KMS: %w", err)
	}

	// decode json and base64 to get encrypted DEK
	var resContent model.EncryptResponse200
	err = json.NewDecoder(res.Body).Decode(&resContent)
	if err != nil {
		return nil, fmt.Errorf("could not decode KMS response with wrapped DEK: %w", err)
	}
	wrappedDEK, err := base64.StdEncoding.DecodeString(resContent.CiphertextB64)
	if err != nil {
		return nil, fmt.Errorf("could not decode base64-encoded wrapped DEK: %w", err)
	}

	return wrappedDEK, nil
}

// Decrypt implements the tink.AEAD interface for decryption.
func (rka *RemoteKEKAEAD) Decrypt(wrappedDEK []byte, aad []byte) ([]byte, error) {
	// prepare and send request
	reqURL := rka.baseURL.JoinPath("/v1/decrypt").String()
	reqBody, _ := json.Marshal(model.DecryptReqBody{
		CiphertextB64: base64.StdEncoding.EncodeToString(wrappedDEK),
	})
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, fmt.Errorf("could not create HTTP request to KMS: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send HTTP request to KMS: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not wrap DEK within KMS: %w", err)
	}

	// decode json and base64 to get DEK
	var resContent model.DecryptResponse200
	err = json.NewDecoder(res.Body).Decode(&resContent)
	if err != nil {
		return nil, fmt.Errorf("could not decode KMS response with wrapped DEK: %w", err)
	}
	dek, err := base64.StdEncoding.DecodeString(resContent.Plaintext)
	if err != nil {
		return nil, fmt.Errorf("could not decode base64-encoded wrapped DEK: %w", err)
	}

	return dek, nil
}
