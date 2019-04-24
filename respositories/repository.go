package respositories

import (
	"database/sql"
	"errors"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	_ "github.com/go-sql-driver/mysql"
)

// NoteRepository - Note repository abstraction
type NoteRepository struct {
	DB *sql.DB
}

// NewNoteRepository - create new object
func NewNoteRepository(db *sql.DB) NoteRepository {
	return NoteRepository{DB: db}
}

// CreateNote - creates new note
func (ur *NoteRepository) CreateNote(note model.Note) (int, error) {
	return -1, errors.New("Not implemented")
}

// UpdateNote - updates note by ID
func (ur *NoteRepository) UpdateNote(id int, note model.Note) (bool, error) {
	return false, errors.New("Not implemented")
}

// DeleteNote - deletes note by ID
func (ur *NoteRepository) DeleteNote(id int) (bool, error) {
	return false, errors.New("Not implemented")
}

// GetNote - gets note by ID
func (ur *NoteRepository) GetNote(id int) (model.Note, error) {
	return model.Note{}, errors.New("Not implemented")
}

// GetNotes - gets notes
func (ur *NoteRepository) GetNotes(start, count int) ([]model.Note, error) {
	return nil, errors.New("Not implemented")
}
