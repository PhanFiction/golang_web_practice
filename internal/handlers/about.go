package handlers

import (
	"goweb_exercise/internal/types"
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").
		ParseFiles(
			"templates/layout.html",
			"templates/about.html",
		))

	data := types.PageData{
		TabTitle:  "About Page",
		PageTitle: "About",
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)
}
