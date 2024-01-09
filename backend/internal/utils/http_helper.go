package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, data interface{}, statusCode ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
