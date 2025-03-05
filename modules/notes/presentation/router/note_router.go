package router

import (
	"note-tracker/commons/config"
	"note-tracker/modules/notes/presentation/controller"

	"github.com/gorilla/mux"
)

// RegisterNoteRoutes sets up API routes
func RegisterNoteRoutes(router *mux.Router, noteController *controller.NoteController) {
	noteRouter := router.PathPrefix(config.NotesAPI).Subrouter()

	// Register routes for different HTTP methods
	noteRouter.HandleFunc("", noteController.HandleCreateNote).Methods("POST")
	noteRouter.HandleFunc("", noteController.HandleUpdateNote).Methods("PUT")
	noteRouter.HandleFunc("/{id}", noteController.HandleDeleteNote).Methods("DELETE")
	noteRouter.HandleFunc("", noteController.HandleGetAllNotes).Methods("GET")
}
