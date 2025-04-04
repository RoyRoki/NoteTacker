package datasource

import (
	"context"
	"database/sql"
	"fmt"
	"note-tracker/modules/notes/data/model"
	"time"
)

type NoteDataSource interface {
	CreateNote(ctx context.Context, note *model.Note) (*model.Note, error)
	DeleteNote(ctx context.Context, id int) error
	GetAllNotes(ctx context.Context) ([]model.Note, error)
	UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error)
}

type noteDataSourceImpl struct {
	db *sql.DB
}

func NewNoteDataSource(db *sql.DB) NoteDataSource {
	return &noteDataSourceImpl{db: db}
}

func (ds *noteDataSourceImpl) CreateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	if note.Date == "" {
		note.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	query := `INSERT INTO notes (title, note, date) VALUES (?, ?, ?)`
	result, err := ds.db.ExecContext(ctx, query, note.Title, note.Note, note.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to insert note: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	note.ID = int(id)
	return note, nil
}

func (ds *noteDataSourceImpl) DeleteNote(ctx context.Context, id int) error {
	query := `DELETE FROM notes WHERE id = ?`
	result, err := ds.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("note with ID %d not found", id)
	}

	return nil
}

func (ds *noteDataSourceImpl) GetAllNotes(ctx context.Context) ([]model.Note, error) {
	query := `SELECT id, title, note, date FROM notes`
	rows, err := ds.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query notes: %w", err)
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Note, &note.Date); err != nil {
			return nil, fmt.Errorf("failed to scan note: %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating notes: %w", err)
	}

	return notes, nil
}

func (ds *noteDataSourceImpl) UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	if note.Date == "" {
		note.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	query := `UPDATE notes SET title = ?, note = ?, date = ? WHERE id = ?`
	_, err := ds.db.ExecContext(ctx, query, note.Title, note.Note, note.Date, note.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update note: %w", err)
	}

	return note, nil
}
