package database

import (
	"database/sql"
	"log"
	"note-tracker/commons/logger"

	_ "github.com/mattn/go-sqlite3"
)

// InitSQLiteDB initializes the SQLight connection.
func InitSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		logger.Error("Failed to connected to SQLite")
		return nil, err
	}

	// Validate the database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping SQLite database: " + err.Error())
		db.Close() // Close connection on failure
		return nil, err
	}

	logger.Info("Database connection established successfully.")
	return db, nil
}

// RunMigrations creates tables if they do not exist
func RunMigrations(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		note TEXT,
		date TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
		return err
	}

	logger.Info("SQLite database migrations applied successfully.")
	return nil
}
