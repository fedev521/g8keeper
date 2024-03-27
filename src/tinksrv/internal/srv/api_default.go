package srv

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/internal/crypt"
	"github.com/fedev521/g8keeper/tinksrv/internal/log"
	"github.com/fedev521/g8keeper/tinksrv/pkg/model"
)

func PostEncrypt(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 64<<10))
	dec.DisallowUnknownFields()

	// decode request
	var body model.EncryptReqBody
	err := dec.Decode(&body)
	if err != nil {
		log.Error(err.Error())
		sendBadRequestError(w, r, "Bad request")
		return
	}

	// encrypt plaintext
	ciphertext, err := crypt.EnvelopeEncrypt([]byte(body.Plaintext), []byte{})
	if err != nil {
		log.Error(err.Error())
		sendUnexpectedServerError(w, r)
		return
	}

	// create response payload with ciphertext
	payload, _ := json.Marshal(model.EncryptResponse200{
		CiphertextB64: base64.StdEncoding.EncodeToString(ciphertext),
	})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func PostDecrypt(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	dec.DisallowUnknownFields()

	// decode request
	var body model.DecryptReqBody
	err := dec.Decode(&body)
	if err != nil {
		log.Error(err.Error())
		sendBadRequestError(w, r, "Bad request")
		return
	}

	// decode and decrypt ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(body.CiphertextB64)
	if err != nil {
		log.Error(err.Error())
		sendBadRequestError(w, r, "Bad request")
		return
	}
	plaintext, err := crypt.EnvelopeDecrypt(ciphertext, []byte{})
	if err != nil {
		log.Error(err.Error())
		sendUnexpectedServerError(w, r)
		return
	}

	// create response payload with plaintext
	payload, _ := json.Marshal(model.DecryptResponse200{
		Plaintext: string(plaintext),
	})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
