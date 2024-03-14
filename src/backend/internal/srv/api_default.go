package srv

import (
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/backend/internal/log"
)

func ListPasswords(w http.ResponseWriter, r *http.Request) {
	// create response payload password information (but not the actual secret)
	payload, err := json.Marshal(ListPasswordsResponse200{
		PasswordMetadata: []PasswordMetadata{
			{Name: "pass1"},
			{Name: "pass2"},
		},
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
