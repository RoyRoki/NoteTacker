package repository_impl

import (
	"context"
	"note-tracker/modules/notes/data/datasource"
	"note-tracker/modules/notes/data/model"
	"note-tracker/modules/notes/domain/repository"
)

// NoteRepositoryImpl implements the NoteRepository interface.
type NoteRepositoryImpl struct {
	ds datasource.NoteDataSource
}

// NewNoteRepository creates a new NoteRepositoryImpl
func NewNoteRepository(ds datasource.NoteDataSource) repository.NoteRepository {
	return &NoteRepositoryImpl{ds: ds}
}

// CreateNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) CreateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	return n.ds.CreateNote(ctx, note)
}

// DeleteNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) DeleteNote(ctx context.Context, id int) error {
	return n.ds.DeleteNote(ctx, id)
}

// GetAllNotes implements repository.NoteRepository.
func (n *NoteRepositoryImpl) GetAllNotes(ctx context.Context) ([]model.Note, error) {
	return n.ds.GetAllNotes(ctx)
}

// UpdateNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	return n.ds.UpdateNote(ctx, note)
}
