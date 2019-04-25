package respositories

import (
	"os"
	"testing"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	_ "github.com/go-sql-driver/mysql"
)

var noteRepository NoteRepository
var sampleNote = model.Note{ID: 1, Title: "sample title", Body: "sample body"}

func TestMain(m *testing.M) {
	noteRepository = NoteRepository{DB: OpenDBConnection("gonotes", "1Q2w3e4r", "gonotes")}

	//noteRepository.EnsureNotesTableExists()
	prepareNotes(noteRepository)
	code := m.Run()
	noteRepository.ClearNotesTable()

	os.Exit(code)

}

func TestNoteRepository_CreateNote(t *testing.T) {
	tests := []struct {
		name    string
		note    model.Note
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", sampleNote, 2, false},
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
		id      int
		note    model.Note
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"happy", 1, sampleNote, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.UpdateNote(tt.id, tt.note)
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
		{"happy", 1, sampleNote, false},
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
		{"happy", 0, 10, []model.Note{sampleNote}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := noteRepository.GetNotes(tt.start, tt.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteRepository.GetNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compareNotes(tt.want, got, t)
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

func prepareNotes(nr NoteRepository) {
	nr.CreateNote(sampleNote)
}
func compareNotes(expecteds, actuals []model.Note, t *testing.T) {
	for i, expected := range expecteds {
		compareNote(expected, actuals[i], t)
	}
}

func compareNote(expected, actual model.Note, t *testing.T) {
	if expected.ID != actual.ID {
		t.Errorf("Expected ID '%d'. Got '%d'\n", expected.ID, actual.ID)
	}
	if expected.Title != actual.Title {
		t.Errorf("Expected Title '%s'. Got '%s'\n", expected.Title, actual.Title)
	}
	if expected.Body != actual.Body {
		t.Errorf("Expected Body '%s'. Got '%s'\n", expected.Body, actual.Body)
	}
}
