package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/login", PostLogin).Methods("POST")
	return router
}
