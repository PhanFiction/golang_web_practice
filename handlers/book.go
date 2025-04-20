package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func BookPageHandler(w http.ResponseWriter, r *http.Request) {
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
}
