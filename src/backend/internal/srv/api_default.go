package srv

import (
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
	"github.com/fedev521/g8keeper/backend/internal/store"
	"github.com/fedev521/g8keeper/backend/internal/types"
)

func ListPasswordsHF(pk store.PasswordKeeper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metadata, err := pk.ListMetadata()
		if err != nil {
			log.Error(err.Error())
			sendUnexpectedServerError(w, r)
			return
		}

		// create response payload password information (but not the actual secret)
		payload, err := json.Marshal(ListPasswordsResponse200{
			PasswordMetadata: metadata,
		})
		if err != nil {
			log.Error(err.Error())
			sendUnexpectedServerError(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}

func PostPasswordHF(pk store.PasswordKeeper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
		dec.DisallowUnknownFields()

		// decode request
		var body CreatePasswordReqBody
		err := dec.Decode(&body)
		if err != nil {
			log.Error(err.Error())
			sendBadRequestError(w, r, "Bad request")
			return
		}

		password := types.Password{
			Secret: body.Secret,
			Metadata: types.PasswordMetadata{
				Name: body.Name,
			},
		}

		err = pk.Store(password, password.Metadata.Name)
		if err != nil {
			log.Error(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
