package repository

import (
	"context"
	"note-tracker/modules/notes/data/model"
)

// NoteRepository defines the operations
type NoteRepository interface {
	// CreateNote insert a new note
	CreateNote(ctx context.Context, note *model.Note) (*model.Note, error)

	// UpdateNote update a note
	UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error)

	// DeleteNote remove a note
	DeleteNote(ctx context.Context, id int) error

	// GetAllNotes retrieves all notes
	GetAllNotes(ctx context.Context) ([]model.Note, error)
}
