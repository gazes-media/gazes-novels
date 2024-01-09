package api

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
