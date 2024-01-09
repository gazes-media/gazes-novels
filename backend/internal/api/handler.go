package api

import (
	"encoding/json"
	"net/http"

	"github.com/gazes-media/gazes-novels/internal/utils"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var userLogin utils.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid request body",
		}, http.StatusBadRequest)
		return
	}
	token, err := utils.AuthPostLogin(userLogin.Email, userLogin.Password)

	if err != nil {
		utils.RespondJSON(w, map[string]string{
			"error": "invalid credentials",
		}, http.StatusUnauthorized)
		return
	}

	utils.RespondJSON(w, map[string]string{
		"token": *token,
	})
}
