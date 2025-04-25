package handlers

import (
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
	tmpl := template.Must(template.ParseFiles("book.html"))
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	pageNumber, err := strconv.Atoi(page)

	if err != nil {
		// handle the error, maybe default to page 1 or return a 400 error
		pageNumber = 1
	}

	data := BookPageData{
		PageTitle:  title,
		PageNumber: pageNumber,
	}

	// fmt.Fprintf(w, `
	// 	<html>
	// 		<body>
	// 			<h1>Welcome to my website!</h1>
	// 			<p>Title is %s </p>
	// 			<p>Page is %s</p>
	// 			<img src="/static/img.jpeg" alt="A cute cat" style="max-width: 300px;" />
	// 		</body>
	// 	</html>
	// `, title, page)

	tmpl.Execute(w, data)
}
