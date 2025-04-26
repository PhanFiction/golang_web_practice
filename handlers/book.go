package handlers

import (
	"goweb_exercise/types"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BookPageData struct {
	PageTitle  string
	PageNumber int
}

func BookPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").
		ParseFiles(
			"templates/layout.html",
			"templates/book.html",
		))

	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		// handle the error, maybe default to page 1 or return a 400 error
		pageNumber = 1
	}

	data := types.PageData{
		TabTitle:  "Book Page",
		PageTitle: title,
		BookPage:  pageNumber,
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)
}
