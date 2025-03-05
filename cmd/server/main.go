package main

import (
	"fmt"
	"log"
	"net/http"
	"note-tracker/commons/database"
	"note-tracker/commons/logger"
	"note-tracker/modules/notes/data/repository_impl"
	"note-tracker/modules/notes/domain/usecase"
	"note-tracker/modules/notes/presentation/controller"
	"note-tracker/modules/notes/presentation/router"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
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

	// Initialize repository, use_case, and controller
	notesRepo := repository_impl.NewNoteRepository(db)
	notesUsecase := usecase.NewNoteUsecase(notesRepo)
	notesController := controller.NewNoteController(notesUsecase)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register routes with the controller using Gorilla Mux
	router.RegisterNoteRoutes(r, notesController)

	// Start the HTTP server
	fmt.Println("Server is starting on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server error:", err)
	}
}
