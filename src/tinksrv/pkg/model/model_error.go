package model

type ErrorResponse struct {
	Message string `json:"error,omitempty"` // Description of the error.
}
