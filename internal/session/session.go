package session

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// Global var to be used
// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func SecretHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.New("layout.html").
			ParseFiles(
				"templates/layout.html",
				"templates/nav.html",
				"templates/login.html",
			))

		tmpl.Execute(w, nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		session, _ := Store.Get(r, "session")
		// Authentication goes here
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Middleware to check if user is authenticated
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
