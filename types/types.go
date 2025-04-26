package types

type PageData struct {
	PageTitle string
	TabTitle  string
	BookPage  int
}

type BookData struct {
	Title       string
	Author      string
	Pages       int
	Publisher   string
	ISBN        string
	Description string
	PublishedAt string
}
