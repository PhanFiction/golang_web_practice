package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, `
			<html>
					<body>
							<h1>Welcome to my website!</h1>
							<p>Title is %s </p>
							<p>Page is %s</p>
							<img src="/static/img.jpeg" alt="A cute cat" style="max-width: 300px;" />
					</body>
			</html>
		`, title, page)
	})

	// Serving static files using gorilla/mux
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Default way to serve static files
	// fs := http.FileServer(http.Dir("static/"))                // point to where static directory is located at
	// http.Handle("/static/", http.StripPrefix("/static/", fs)) // Strip away part of the url path and serve the files from here

	http.ListenAndServe(":8080", r)
}
