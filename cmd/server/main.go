package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"note-tracker/commons/database"
	"note-tracker/commons/logger"
	"note-tracker/commons/middleware"
	"note-tracker/modules/notes/di"
)

func main() {
	// Initialize logger early.
	logger.InitLogger()

	// Initialize sqlite3 connection.
	db, err := database.InitSQLiteDB()
	if err != nil {
		log.Fatalf("Database initialization failed: %s", err)
	}
	defer db.Close()

	// RunMigrations creates tables.
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register routes with the controller using Gorilla Mux
	di.RegisterNoteModule(db, r)

	corsWrapped := middleware.CORS(r)

	// Start the HTTP server with custom CORS middleware
	fmt.Println("Server is starting on :8080...")
	if err := http.ListenAndServe(":8080", corsWrapped); err != nil {
		fmt.Println("Server error:", err)
	}
}
