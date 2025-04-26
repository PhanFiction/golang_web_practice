package main

import (
	"goweb_exercise/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.SetupRoutes(r)
	http.ListenAndServe(":8080", r)
}
