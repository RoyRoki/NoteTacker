package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// SQLite connection
var DB *sql.DB

// InitDB initializes the database connection and creates the notes table if necessary
func InitDB() {
	var err error
	// Open the SQLite database
	DB, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	// Create the notes table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		note TEXT,
		date TEXT
	)`)
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
	}

	log.Println("Database initialized successfully.")
}
