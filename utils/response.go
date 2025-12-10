package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"twitter-clone/pkg/domain"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fallback := map[string]string{"error": "failed to encode JSON response"}

		_ = json.NewEncoder(w).Encode(fallback)
	}
}

func DecodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func Error(w http.ResponseWriter, code int, message string, details ...any) {
	if len(details) > 0 {
		log.Printf("[ERROR] %s | details: %v", message, details)
	} else {
		log.Printf("[ERROR] %s", message)
	}

	RespondJSON(w, code, ErrorResponse{Error: message})
}

func HandleAPIError(w http.ResponseWriter, err error) {
	var code int
	var message string

	if getter, ok := err.(domain.HTTPStatusGetter); ok {
		code = getter.Status()
		message = getter.Error()
	} else {
		switch {
		case errors.Is(err, strconv.ErrRange) || errors.Is(err, strconv.ErrSyntax):
			code = http.StatusBadRequest
			message = "invalid format or range for an ID"

		default:
			code = http.StatusInternalServerError
			message = "server error"
			Error(w, code, message, err)
			return
		}
	}

	Error(w, code, message)
}
