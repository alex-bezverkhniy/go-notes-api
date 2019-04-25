package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	"github.com/alex-bezverkhniy/go-notes-api/respositories"
)

var a App
var nr respositories.NoteRepository

func TestMain(m *testing.M) {
	a = NewApp("gonotes", "1Q2w3e4r", "gonotes", "", "3306")

	// a.NoteRepository.EnsureNotesTableExists()
	prepareNotes(a.NoteRepository)
	code := m.Run()
	a.NoteRepository.ClearNotesTable()

	os.Exit(code)
}

func TestNotes(t *testing.T) {
	req, _ := http.NewRequest("GET", APIBasePath+"/notes", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	if body := res.Body.String(); body == "[]" {
		t.Errorf("Expected NON empty payload (array). Got %s", body)
	}

	expected := []model.Note{model.Note{ID: 1, Title: "sample title", Body: "sample body"}}
	var actual []model.Note
	err := json.Unmarshal(res.Body.Bytes(), &actual)

	if err != nil || actual == nil {
		t.Errorf("Expected Non empty array. Got %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestGetNonExistenNote(t *testing.T) {
	req, _ := http.NewRequest("GET", APIBasePath+"/notes/42", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)
	var m map[string]string
	json.Unmarshal(res.Body.Bytes(), &m)
	if m["error"] != respositories.NoteNotFoundMsg {
		t.Errorf("Expected the 'error' key of the response to be set to '%s'. Got '%s'", respositories.NoteNotFoundMsg, m["error"])
	}
}

func TestCreateNote(t *testing.T) {
	payload := []byte(`{"title": "sample title", "body": "sample body"}`)

	req, _ := http.NewRequest("POST", APIBasePath+"/notes", bytes.NewBuffer(payload))
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	loc := res.HeaderMap.Get("Location")
	if loc != APIBasePath+"/notes/2" {
		t.Errorf("Expected the '"+APIBasePath+"/notes/2' in header 'Location'. Got '%s'", loc)
	}
}

func TestGetNote(t *testing.T) {
	expected := model.Note{ID: 1, Title: "sample title", Body: "sample body"}
	req, _ := http.NewRequest("GET", APIBasePath+"/notes/1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	var actual model.Note
	json.Unmarshal(res.Body.Bytes(), &actual)
	compareNotes(expected, actual, t)
}

func TestUpdateNote(t *testing.T) {
	expectedNote := model.Note{ID: 1, Title: "sample updated title", Body: "sample updated body"}
	payload, _ := json.Marshal(expectedNote)

	req, _ := http.NewRequest("PUT", APIBasePath+"/notes/1", bytes.NewBuffer(payload))
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	var actual bool
	json.Unmarshal(res.Body.Bytes(), &actual)
	if actual != true {
		t.Errorf("Expected the 'true' but got '%v'", actual)
	}

	req, _ = http.NewRequest("GET", APIBasePath+"/notes/1", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	var actualNote model.Note
	json.Unmarshal(res.Body.Bytes(), &actualNote)
	compareNotes(expectedNote, actualNote, t)
}

func TestUpdateNonExistNote(t *testing.T) {
	expectedNote := model.Note{ID: 1, Title: "sample updated title", Body: "sample updated body"}
	payload, _ := json.Marshal(expectedNote)

	req, _ := http.NewRequest("PUT", APIBasePath+"/notes/42", bytes.NewBuffer(payload))
	res := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)

	req, _ = http.NewRequest("GET", APIBasePath+"/notes/42", nil)
	res = executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func TestDeleteNote(t *testing.T) {

	req, _ := http.NewRequest("DELETE", APIBasePath+"/notes/1", nil)
	res := executeRequest(req)

	checkResponseCode(t, http.StatusOK, res.Code)
	var actual bool
	json.Unmarshal(res.Body.Bytes(), &actual)
	if actual != true {
		t.Errorf("Expected the 'true' but got '%v'", actual)
	}

	req, _ = http.NewRequest("GET", APIBasePath+"/notes/1", nil)
	res = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, res.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func prepareNotes(nr respositories.NoteRepository) {
	nr.CreateNote(model.Note{Title: "sample title", Body: "sample body"})
}
func compareNotes(expected, actual model.Note, t *testing.T) {
	if expected.ID != actual.ID {
		t.Errorf("Expected ID %d. Got %d\n", expected.ID, actual.ID)
	}
	if expected.Title != actual.Title {
		t.Errorf("Expected Title %s. Got %s\n", expected.Title, actual.Title)
	}
	if expected.Body != actual.Body {
		t.Errorf("Expected Body %s. Got %s\n", expected.Body, actual.Body)
	}
}
