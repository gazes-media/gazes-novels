package api

import (
	"context"
	"net/http"

	"github.com/gazes-media/gazes-novels/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			utils.RespondJSON(w, map[string]string{
				"error": "missing Authorization header",
			}, http.StatusBadRequest)
			return
		}

		if authHeader[:7] != "Bearer " {
			utils.RespondJSON(w, map[string]string{
				"error": "invalid Authorization header",
			}, http.StatusBadRequest)
			return
		}

		err, user := utils.AuthGetUserMe(authHeader[7:])
		if err != nil {
			utils.RespondJSON(w, map[string]string{
				"error": "invalid token",
			}, http.StatusUnauthorized)
			return
		}
		const userKey contextKey = "user"
		ctx := context.WithValue(r.Context(), userKey, user)
		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey string
