package respositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	_ "github.com/go-sql-driver/mysql"
)

var noteRepository NoteRepository

func TestMain(m *testing.M) {
	noteRepository = NoteRepository{DB: openConnection("gonotes", "1Q2w3e4r", "gonotes")}

	ensureNotesTableExists()
	code := m.Run()
	clearNotesTable()

	os.Exit(code)

}

func ensureNotesTableExists() {
	if _, err := noteRepository.DB.Exec(model.NoteDDL); err != nil {
		log.Print(err)
	} else {
		log.Fatal("Table Note is not exist!")
	}
}

func clearNotesTable() {
	noteRepository.DB.Exec("DELETE FROM notes")
	noteRepository.DB.Exec("ALTER TABLE notes AUTO_INCREMENT = 1")
}

func openConnection(user, password, dbname string) *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestNoteRepository_CreateNote(t *testing.T) {
	tests := []struct {
		name    string
		note    model.Note
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", model.Note{0, "sample note title", "sample note body"}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.CreateNote(tt.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.CreateNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NoteRepository.CreateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_UpdateNote(t *testing.T) {
	tests := []struct {
		name    string
		note    model.Note
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", model.Note{0, "sample note title", "sample note body"}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.UpdateNote(1, tt.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.UpdateNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NoteRepository.UpdateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_GetNote(t *testing.T) {
	tests := []struct {
		name    string
		noteId  int
		want    model.Note
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", 1, model.Note{1, "sample note title", "sample note body"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.GetNote(tt.noteId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.GetNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NoteRepository.GetNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_GetNotes(t *testing.T) {
	tests := []struct {
		name    string
		start   int
		count   int
		want    []model.Note
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", 1, 1, []model.Note{model.Note{1, "sample note title", "sample note body"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.GetNotes(tt.start, tt.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.GetNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteRepository.GetNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteRepository_DeleteNote(t *testing.T) {
	tests := []struct {
		name    string
		noteId  int
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", 1, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.DeleteNote(tt.noteId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.DeleteNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NoteRepository.DeleteNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
