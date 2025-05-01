package routes

import (
	"goweb_exercise/internal/handlers"
	"goweb_exercise/internal/middleware"
	"goweb_exercise/internal/session"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/about", middleware.Chain(handlers.AboutHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/books/{title}/page/{page}", middleware.Chain(handlers.BookPageHandler, middleware.Logging())).Methods("GET")
	r.HandleFunc("/create-book", middleware.Chain(handlers.CreateBookHandler, middleware.Logging())).Methods("POST")
	r.HandleFunc("/dashboard", middleware.Chain(session.AuthMiddleware(handlers.DashboardHandler), middleware.Logging()))
	r.HandleFunc("/", middleware.Chain(handlers.HomeHandler, middleware.Logging())).Methods("GET")

	// Auth
	r.HandleFunc("/login", middleware.Chain(session.LoginHandler, middleware.Logging()))
	r.HandleFunc("/logout", middleware.Chain(session.LogoutHandler, middleware.Logging()))
	r.HandleFunc("/secret", middleware.Chain(session.SecretHandler, middleware.Logging()))

	// Serve CSS, JS, and images from static dir
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
