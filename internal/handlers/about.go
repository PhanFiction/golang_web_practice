package handlers

import (
	"goweb_exercise/internal/session"
	"goweb_exercise/internal/types"
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").
		ParseFiles(
			"templates/layout.html",
			"templates/nav.html",
			"templates/about.html",
		))

	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "About Page",
		PageTitle:     "About",
		Authenticated: ok && auth,
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)
}
