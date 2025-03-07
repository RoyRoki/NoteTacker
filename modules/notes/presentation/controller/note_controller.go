package controller

import (
	"encoding/json"
	"net/http"
	"note-tracker/commons/strings"
	"note-tracker/modules/notes/data/model"
	"note-tracker/modules/notes/domain/usecase"
	"note-tracker/modules/notes/presentation/dto"
	"strconv"

	"github.com/gorilla/mux"
)

// NoteController handles HTTP requests related to notes
type NoteController struct {
	NoteUsecase *usecase.NoteUsecase
}

// NewNoteController initializes a controller with a usecase
func NewNoteController(notesUsecase *usecase.NoteUsecase) *NoteController {
	return &NoteController{NoteUsecase: notesUsecase}
}

// HandleCreateNote processes note creation requests
func (c *NoteController) HandleCreateNote(w http.ResponseWriter, r *http.Request) {
	var note_dto dto.CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&note_dto); err != nil {
		http.Error(w, strings.InvalidCreateNoteRequest, http.StatusBadRequest)
		return
	}

	noteModel, err := note_dto.ToModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdNote, err := c.NoteUsecase.CreateNote(r.Context(), noteModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdNote)
}

// HandleUpdateNote processes note update requests
func (c *NoteController) HandleUpdateNote(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, strings.InvalidRequestData, http.StatusBadRequest)
		return
	}

	updatedNote, err := c.NoteUsecase.UpdateNote(r.Context(), &note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(updatedNote)
}

// HandleDeleteNote handles note deletion requests
func (c *NoteController) HandleDeleteNote(w http.ResponseWriter, r *http.Request) {
	// Extract 'id' from the URL path
	vars := mux.Vars(r)
	idParam := vars["id"] // Extracts the 'id' from the URL path

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "invalid note ID", http.StatusBadRequest)
		return
	}

	err = c.NoteUsecase.DeleteNote(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HandleGetAllNotes handles fetching all notes
func (c *NoteController) HandleGetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := c.NoteUsecase.GetAllNotes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(notes)
}
