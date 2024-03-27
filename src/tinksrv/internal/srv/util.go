package srv

import (
	"encoding/json"
	"net/http"

	"github.com/fedev521/g8keeper/tinksrv/pkg/model"
)

func sendError(w http.ResponseWriter, _ *http.Request, code int, e *model.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}

func sendBadRequestError(w http.ResponseWriter, r *http.Request, msg string) {
	sendError(w, r, http.StatusBadRequest, &model.ErrorResponse{
		Message: msg,
	})
}

func sendUnexpectedServerError(w http.ResponseWriter, r *http.Request) {
	sendError(w, r, http.StatusInternalServerError, &model.ErrorResponse{
		Message: "Something went wrong",
	})
}
