package usecase

import (
	"context"
	"errors"
	"note-tracker/modules/notes/data/model"
	"note-tracker/modules/notes/domain/repository"
)

// NoteUsecase defines business logic for notes
type NoteUsecase struct {
	NoteRepo repository.NoteRepository
}

// NewNoteUsecase initializes the usecase with a repository
func NewNoteUsecase(notesRepo repository.NoteRepository) *NoteUsecase {
	return &NoteUsecase{NoteRepo: notesRepo}
}

// CreateNote applies business rules before saving a note
func (u *NoteUsecase) CreateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	// Validate input data
	if note.Title == "" || note.Note == "" {
		return nil, errors.New("tittle and note cannot be empty")
	}

	// Call the repository to create new note
	return u.NoteRepo.CreateNote(ctx, note)
}

// UpdateNote applies business rules before updating a note
func (u *NoteUsecase) UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	// Validate input data
	if note.ID == 0 {
		return nil, errors.New("id cannot be empty")
	}
	if note.Title == "" || note.Note == "" {
		return nil, errors.New("tittle and note cannot be empty")
	}

	// Call the repository to update the note
	return u.NoteRepo.UpdateNote(ctx, note)
}

// DeleteNote handles business logic for deleting a note
func (u *NoteUsecase) DeleteNote(ctx context.Context, id int) error {
	// Validate input data
	if id == 0 {
		return errors.New("note ID is required")
	}

	// Call repository to delete the note
	return u.NoteRepo.DeleteNote(ctx, id)
}

// GetAllNotes fetches all notes from the repository
func (u *NoteUsecase) GetAllNotes(ctx context.Context) ([]model.Note, error) {
	return u.NoteRepo.GetAllNotes(ctx)
}
