package srv

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/internal/kms"
	"github.com/fedev521/g8keeper/tinksrv/internal/log"
	"github.com/fedev521/g8keeper/tinksrv/pkg/model"
)

func PostEncryptHF(kekManager *kms.KEKManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// TODO handle _ errors

		// encrypt plaintext
		primitive, _ := kekManager.GetAEAD()
		pt := []byte(body.PlaintextB64)
		aad := []byte("aad")
		ciphertext, err := primitive.Encrypt(pt, aad)
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
}

func PostDecryptHF(kekManager *kms.KEKManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 64<<10))
		dec.DisallowUnknownFields()

		// decode request
		var body model.DecryptReqBody
		err := dec.Decode(&body)
		if err != nil {
			log.Error(err.Error())
			sendBadRequestError(w, r, "Bad request")
			return
		}

		// base64-decode and decrypt ct
		ct, err := base64.StdEncoding.DecodeString(body.CiphertextB64)
		aad := []byte("aad")
		if err != nil {
			log.Error(err.Error())
			sendBadRequestError(w, r, "Bad request")
			return
		}

		primitive, err := kekManager.GetAEAD()
		if err != nil {
			log.Error(err.Error())
			sendUnexpectedServerError(w, r)
			return
		}

		plaintext, err := primitive.Decrypt(ct, aad)
		if err != nil {
			log.Error(err.Error())
			sendUnexpectedServerError(w, r)
			return
		}

		// TODO handle _ error

		// create response payload with plaintext
		payload, _ := json.Marshal(model.DecryptResponse200{
			Plaintext: string(plaintext),
		})

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}
