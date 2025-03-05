package repository_impl

import (
	"context"
	"database/sql"
	"fmt"
	"note-tracker/modules/notes/data/model"
	"note-tracker/modules/notes/domain/repository"
	"time"
)

// NoteRepositoryImpl implements the NoteRepository interface.
type NoteRepositoryImpl struct {
	DB *sql.DB
}

// NewNoteRepository creates a new NoteRepositoryImpl
func NewNoteRepository(db *sql.DB) repository.NoteRepository {
	return &NoteRepositoryImpl{DB: db}
}

// CreateNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) CreateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	// Check if the date is provided, if not, set the current date
	if note.Date == "" {
		// Set the date to the current time in the format "YYYY-MM-DD HH:MM:SS"
		note.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	// Insert a new note into the notes table
	insertSQL := `INSERT INTO notes (title, note, date) VALUES (?, ?, ?)`
	result, err := n.DB.ExecContext(ctx, insertSQL, note.Title, note.Note, note.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to execute insert query: %w", err)
	}

	// Get the last inserted ID and update the note with it
	id, err := result.LastInsertId()
	if err != nil {
		// Return error if unable to retrieve the inserted ID
		return nil, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	// Ensure the ID is valid (it should not be zero)
	if id <= 0 {
		return nil, fmt.Errorf("invalid last insert ID: %d", id)
	}

	// Set the ID for the note and return the created note
	note.ID = int(id)
	return note, nil
}

// DeleteNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) DeleteNote(ctx context.Context, id int) error {
	// Prepare the SQL DELETE statement
	deleteSQL := `DELETE FROM notes WHERE id = ?`
	result, err := n.DB.ExecContext(ctx, deleteSQL, id)
	if err != nil {
		return fmt.Errorf("failed to delete note: %v", err)
	}

	// Check if any row was affected (note exists)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("note with ID %d not found", id)
	}

	return nil
}

// GetAllNotes implements repository.NoteRepository.
func (n *NoteRepositoryImpl) GetAllNotes(ctx context.Context) ([]model.Note, error) {
	// Query to fetch all notes
	querySQL := `SELECT id, title, note, date FROM notes`
	rows, err := n.DB.QueryContext(ctx, querySQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get all notes: %v", err)
	}
	defer rows.Close()

	var notes []model.Note
	// Iterate through the rows and add each note to the slice
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Note, &note.Date); err != nil {
			return nil, fmt.Errorf("failed to scan note: %v", err)
		}
		notes = append(notes, note)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return notes, nil
}

// UpdateNote implements repository.NoteRepository.
func (n *NoteRepositoryImpl) UpdateNote(ctx context.Context, note *model.Note) (*model.Note, error) {
	// Check if the date is provided, if not, set the current date
	if note.Date == "" {
		// Set the date to the current time in the format "YYYY-MM-DD HH:MM:SS"
		note.Date = time.Now().Format("2006-01-02 15:04:05")
	}

	// Prepare the SQL UPDATE statement
	updateSQL := `UPDATE notes SET title = ?, note = ?, date = ? WHERE id = ?`
	_, err := n.DB.ExecContext(ctx, updateSQL, note.Title, note.Note, note.Date, note.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update note: %v", err)
	}

	// Return the updated note
	return note, nil
}
