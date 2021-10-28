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

func badRequest(msg string) httpError {
	return httpError{StatusCode: http.StatusBadRequest, Type: "Bad Request", Message: msg}
}

func notFound(msg string) httpError {
	return httpError{StatusCode: http.StatusNotFound, Type: "Unknown", Message: msg}
}

func unauthoriedRequest(msg string) httpError {
	return httpError{StatusCode: http.StatusUnauthorized, Type: "Unauthorized", Message: msg}
}

func internalServerError(msg string) httpError {
	return httpError{StatusCode: http.StatusInternalServerError, Type: "InternalServerError", Message: msg}
}

func Error(errorType string, msg string) httpError {
	var err httpError

	switch errorType {
	case "Unauthorized":
		err = unauthoriedRequest(msg)
	case "NotFound":
		err = notFound(msg)
	case "BadRequest":
		err = badRequest(msg)
	default:
		err = internalServerError(msg)
	}

	return err
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
