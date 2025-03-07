package dto

import (
	"note-tracker/modules/notes/data/model"
	"time"
)

// CreateNoteRequest defines the expected payload for creating a note.
type CreateNoteRequest struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

// ToModel converts CreateNoteRequest into a model.Note
func (req *CreateNoteRequest) ToModel() (*model.Note, error) {
	now := time.Now().Format("2006-01-02 15:04:05")

	return &model.Note{
		ID:    0,
		Title: req.Title,
		Note:  req.Note,
		Date:  now,
	}, nil
}
