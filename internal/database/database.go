package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB // Global DB connection

func init() {
	err := godotenv.Load()
	fmt.Println("Init has been called")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Optional: ping to ensure connection is valid
	if err := DB.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}
}

func Hello() {
	fmt.Println("Hello World")
}

func GetUsersTable(db *sql.DB) []string {
	query := `SELECT username FROM usernames;`

	rows, _ := db.Query(query)

	var usernames []string

	for rows.Next() {
		var username string
		rows.Scan(&username)
		usernames = append(usernames, username)
	}

	fmt.Println(usernames)
	return usernames
}

func CreateUserTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		google_id TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		name TEXT,
		avatar_url TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	fmt.Println("Users table created or already exists.")
}
