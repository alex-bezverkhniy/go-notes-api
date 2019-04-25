package respositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	_ "github.com/go-sql-driver/mysql"
)

// NoteNotFoundMsg - error msg
const NoteNotFoundMsg = "note not found"

// NoteRepository - Note repository abstraction
type NoteRepository struct {
	DB *sql.DB
}

// NewNoteRepository - create new object
func NewNoteRepository(db *sql.DB) NoteRepository {
	return NoteRepository{DB: db}
}

// CreateNote - creates new note
func (nr *NoteRepository) CreateNote(note model.Note) (int, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO notes(title, body) VALUES('%s','%s')", note.Title, note.Body)

	if _, err := nr.DB.Exec(sqlQuery); err != nil {
		return -1, err
	}

	err := nr.DB.QueryRow("SELECT LAST_INSERT_ID()").Scan(&note.ID)
	if err != nil {
		return -1, err
	}

	return note.ID, nil
}

// UpdateNote - updates note by ID
func (nr *NoteRepository) UpdateNote(id int, note model.Note) (bool, error) {
	sqlQuery := fmt.Sprintf("UPDATE notes SET title='%s', body='%s' WHERE id=%d", note.Title, note.Body, id)

	if _, err := nr.DB.Exec(sqlQuery); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteNote - deletes note by ID
func (nr *NoteRepository) DeleteNote(id int) (bool, error) {
	sqlQuery := fmt.Sprintf("DELETE FROM notes WHERE id = %d", id)

	if _, err := nr.DB.Exec(sqlQuery); err != nil {
		return false, err
	}
	return true, nil
}

// GetNote - gets note by ID
func (nr *NoteRepository) GetNote(id int) (model.Note, error) {
	sqlQuery := fmt.Sprintf("SELECT id, title, body FROM notes WHERE id = %d", id)

	row := nr.DB.QueryRow(sqlQuery)
	note := model.Note{}
	row.Scan(&note.ID, &note.Title, &note.Body)

	if note.ID != id {
		return model.Note{}, errors.New(NoteNotFoundMsg)
	}

	return note, nil
}

// GetNotes - gets notes
func (nr *NoteRepository) GetNotes(start, count int) ([]model.Note, error) {
	sqlQuery := fmt.Sprintf("SELECT id, title, body FROM notes LIMIT %d OFFSET %d", count, start)

	rows, err := nr.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []model.Note{}

	for rows.Next() {
		var n model.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Body); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	return notes, nil
}

// IsNoteExist - return true if note with ID exists
func (nr *NoteRepository) IsNoteExist(id int) (bool, error) {
	sqlQuery := fmt.Sprintf("SELECT count(id) FROM notes WHERE id=%d", id)
	count := -1

	err := nr.DB.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// OpenDBConnection - opens connection with DB
func OpenDBConnection(user, password, dbname string) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// EnsureNotesTableExists - checks if tabel does NOT exist - creates it
func (nr *NoteRepository) EnsureNotesTableExists() {
	if _, err := nr.DB.Exec(model.NoteDDL); err != nil {
		log.Print(err)
	} else {
		log.Fatal("Table Note is not exist!")
	}
}

// ClearNotesTable - Drop all records from table
func (nr *NoteRepository) ClearNotesTable() {
	nr.DB.Exec("DELETE FROM notes")
	nr.DB.Exec("ALTER TABLE notes AUTO_INCREMENT = 1")
}
