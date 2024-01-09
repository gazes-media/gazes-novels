package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gazes-media/gazes-novels/internal/utils"
	"gorm.io/gorm"
)

const (
	authUrl string     = "https://api.gazes.fr/auth"
	userKey contextKey = "user"
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

		httpClient := &http.Client{}
		req, err := http.NewRequest("POST", authUrl+"/@me", nil)
		if err != nil {
			utils.RespondJSON(w, map[string]string{
				"error": "failed to create request",
			}, http.StatusInternalServerError)
			return
		}

		// send the token in the Authorization header
		req.Header.Set("Authorization", "Bearer "+authHeader[7:])
		res, err := httpClient.Do(req)
		if err != nil {
			utils.RespondJSON(w, map[string]string{
				"error": "failed to send request",
			}, http.StatusInternalServerError)
			return
		}

		if res.StatusCode != http.StatusOK {
			utils.RespondJSON(w, map[string]string{
				"error": "invalid token",
			}, http.StatusUnauthorized)
			return
		}

		// We found the user, deserialize the response body into a User struct
		var user UserMiddleware
		const userKey contextKey = "user"

		if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
			utils.RespondJSON(w, map[string]string{
				"error": "failed to decode response body",
			}, http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, user)
		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type UserMiddleware struct {
	gorm.Model
	Username string `json:"username"`            // Username is the user's username.
	Email    string `gorm:"unique" json:"email"` // Email is the user's email address.
	Password string `json:"password"`            // Password is the user's password.
}

type contextKey string
