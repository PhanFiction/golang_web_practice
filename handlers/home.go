package handlers

import (
	"goweb_exercise/types"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").
		ParseFiles(
			"templates/layout.html",
			"templates/home.html",
		))

	data := types.PageData{
		TabTitle:  "Home Page",
		PageTitle: "Home",
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)
}
