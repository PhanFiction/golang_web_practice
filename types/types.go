package types

import "net/http"

type PageData struct {
	PageTitle string
	TabTitle  string
	BookPage  int
}

type BookDetails struct {
	Title       string
	Author      string
	Pages       int
	Publisher   string
	ISBN        string
	Description string
	PublishedAt string
}

// Type function that takes an http.HandlerFunc and returns another http.HandlerFunc.
type Middleware func(http.HandlerFunc) http.HandlerFunc
