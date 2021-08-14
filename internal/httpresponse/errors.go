package httpresponse

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func BadRequest(msg string) httpError {
	return httpError{StatusCode: http.StatusBadRequest, Type: "Bad Request", Message: msg}
}

func NotFound(msg string) httpError {
	return httpError{StatusCode: http.StatusNotFound, Type: "Unknown", Message: msg}
}

func UnauthoriedRequest(msg string) httpError {
	return httpError{StatusCode: http.StatusUnauthorized, Type: "Unauthorze", Message: msg}
}

func (e httpError) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
