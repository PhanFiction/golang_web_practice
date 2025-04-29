package routes

import (
	"goweb_exercise/internal/handlers"
	"goweb_exercise/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", middleware.Chain(handlers.HomeHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/about", middleware.Chain(handlers.AboutHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/books/{title}/page/{page}", middleware.Chain(handlers.BookPageHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/create-book", middleware.Chain(handlers.CreateBookHandler, middleware.Logging())).Methods("POST")

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
