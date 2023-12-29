package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes a JSON response with the given status code and response body.
// The response body is encoded to JSON. The Content-Type header is set to
// application/json.
func RespondJSON(statusCode int, responseBody interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseBody)
}
