package utils

import (
	"encoding/json"
	"log"
	"net/http"
	handlers "twitter-clone/pkg/types"
)

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

	RespondJSON(w, code, handlers.ErrorResponse{Error: message})
}
