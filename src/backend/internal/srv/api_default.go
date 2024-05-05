package srv

import (
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
	"github.com/fedev521/g8keeper/backend/internal/store"
	"github.com/fedev521/g8keeper/backend/internal/types"
	"github.com/gorilla/mux"
)

func GetPasswordHF(pk store.PasswordKeeper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathParams := mux.Vars(r)
		passwordId, found := pathParams[passwordIdKey]
		if !found {
			log.Error("could not get password ID", map[string]interface{}{
				"key": passwordIdKey,
			})
			sendNotFoundError(w, r, "Password matching given id not found")
			return
		}
		// TODO escape and validate param
		password, err := pk.Retrieve(passwordId)
		if err != nil {
			// TODO handle client and server errors differently
			log.Error(err.Error())
			sendNotFoundError(w, r, "Password not found")
			return
		}

		// create response payload with both password secret and metadata
		payload, err := json.Marshal(GetPasswordResponse200{
			Password: password,
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

func ListPasswordsHF(pk store.PasswordKeeper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metadata, err := pk.ListMetadata()
		if err != nil {
			log.Error(err.Error())
			sendUnexpectedServerError(w, r)
			return
		}

		// create response payload with password information (but no secret)
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
