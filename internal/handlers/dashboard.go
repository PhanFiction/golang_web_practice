package handlers

import (
	"encoding/json"
	"goweb_exercise/internal/session"
	"goweb_exercise/internal/types"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.New("layout.html").
	// 	ParseFiles(
	// 		"templates/layout.html",
	// 		"templates/nav.html",
	// 		"templates/dashboard.html",
	// 	))

	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		Username:      "Alex",
		TabTitle:      "Dashboard",
		Authenticated: ok && auth,
	}

	json.NewEncoder(w).Encode(data)

	//tmpl.ExecuteTemplate(w, "layout.html", data)
}
