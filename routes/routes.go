package routes

import (
	"goweb_exercise/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/about", handlers.AboutHandler).Methods("GET")

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
