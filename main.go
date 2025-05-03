package main

import (
	"goweb_exercise/internal/database"
	"goweb_exercise/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	defer database.DB.Close()
	database.GetUsersTable(database.DB)

	// database.CreateUserTable(db)

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	http.ListenAndServe(":8080", r)
}
