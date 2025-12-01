package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		fallback := map[string]string{"error": "failed to encode JSON response"}

		_ = json.NewEncoder(w).Encode(fallback)
	}
}

func Error(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := map[string]string{"error": message}

	_ = json.NewEncoder(w).Encode(resp)
}
