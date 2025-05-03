package handlers

import (
	"encoding/json"
	"goweb_exercise/internal/session"
	"goweb_exercise/internal/types"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.New("layout.html").
	// 	ParseFiles(
	// 		"templates/layout.html",
	// 		"templates/nav.html",
	// 		"templates/home.html",
	// 	))

	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "Home Page",
		PageTitle:     "Home",
		Authenticated: ok && auth,
	}

	//tmpl.ExecuteTemplate(w, "layout.html", data)

	json.NewEncoder(w).Encode(data)
}
