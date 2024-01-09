package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gazes-media/gazes-novels/internal/api"
)

func main() {
	router := api.NewRouter()

	fmt.Printf("Listening on port %s...\n", os.Getenv("PORT"))
	http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), router)

}
