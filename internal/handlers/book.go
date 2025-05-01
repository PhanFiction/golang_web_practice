package handlers

import (
	"fmt"
	"goweb_exercise/internal/session"
	"goweb_exercise/internal/types"
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
			"templates/nav.html",
			"templates/book.html",
			"templates/book_form.html",
		))

	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		// handle the error, maybe default to page 1 or return a 400 error
		pageNumber = 1
	}

	// Fetch user
	session, _ := session.Store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)

	data := types.PageData{
		TabTitle:      "Book Page",
		PageTitle:     title,
		BookPage:      pageNumber,
		Authenticated: ok && auth,
	}

	tmpl.ExecuteTemplate(w, "layout.html", data)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("book_form.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	pageNumber, err := strconv.Atoi(r.FormValue("pages"))

	if err != nil {
		// handle the error, maybe default to page 1 or return a 400 error
		pageNumber = 1
	}

	details := types.BookDetails{
		Title:       r.FormValue("title"),
		Author:      r.FormValue("author"),
		Pages:       pageNumber,
		Publisher:   r.FormValue("publisher"),
		ISBN:        r.FormValue("isbn"),
		Description: r.FormValue("description"),
		PublishedAt: r.FormValue("publishedat"),
	}

	fmt.Fprintf(w, "Book received: %+v", details)
}
