package srv

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, _ *http.Request, code int, e *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(e)
}

func sendBadRequestError(w http.ResponseWriter, r *http.Request, msg string) {
	sendError(w, r, http.StatusBadRequest, &ErrorResponse{
		Message: msg,
	})
}

func sendUnexpectedServerError(w http.ResponseWriter, r *http.Request) {
	sendError(w, r, http.StatusInternalServerError, &ErrorResponse{
		Message: "Something went wrong",
	})
}

func sendNotFoundError(w http.ResponseWriter, r *http.Request, msg string) {
	sendError(w, r, http.StatusNotFound, &ErrorResponse{
		Message: msg,
	})
}
