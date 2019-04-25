package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alex-bezverkhniy/go-notes-api/model"
	"github.com/alex-bezverkhniy/go-notes-api/respositories"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

// APIBasePath - Base path of API
const APIBasePath = "/api/v1"

// DefaultNotesCount - default count notes on page
const DefaultNotesCount = 10

// App - main app struncture
type App struct {
	Router         *mux.Router
	NoteRepository respositories.NoteRepository
}

// NewApp - init function
func NewApp(user, password, dbname string) App {
	a := App{}
	a.NoteRepository = respositories.NewNoteRepository(respositories.OpenDBConnection(user, password, dbname))

	a.Router = mux.NewRouter()
	a.initializeRouters()

	return a
}

// Run - runs api
func (a *App) Run(address string) {
	log.Println("Listening: ", address)
	log.Fatal(http.ListenAndServe(address, a.Router))
}

func (a *App) initializeRouters() {
	a.Router.HandleFunc(APIBasePath+"/notes", a.getNotes).Methods("GET")
	a.Router.HandleFunc(APIBasePath+"/notes", a.createNote).Methods("POST")
	a.Router.HandleFunc(APIBasePath+"/notes/{id:[0-9]+}", a.getNote).Methods("GET")
	a.Router.HandleFunc(APIBasePath+"/notes/{id:[0-9]+}", a.updateNote).Methods("PUT")
	a.Router.HandleFunc(APIBasePath+"/notes/{id:[0-9]+}", a.deleteNote).Methods("DELETE")
}

func (a *App) getNotes(w http.ResponseWriter, r *http.Request) {
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		count = DefaultNotesCount
	}

	start, err := strconv.Atoi(r.FormValue("start"))
	if err != nil {
		start = 0
	}

	if count > 10 || count < 1 {
		count = DefaultNotesCount
	}
	if start < 0 {
		start = 0
	}

	notes, err := a.NoteRepository.GetNotes(start, count)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, notes)
}

func (a *App) getNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	note, err := a.NoteRepository.GetNote(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Note with ID: '"+strconv.Itoa(id)+"' Not found")
		default:
			status := http.StatusInternalServerError
			if err.Error() == respositories.NoteNotFoundMsg {
				status = http.StatusNotFound
			}
			respondWithError(w, status, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, note)
}

func (a *App) createNote(w http.ResponseWriter, r *http.Request) {
	var n model.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	ID, err := a.NoteRepository.CreateNote(n)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Location", APIBasePath+"/notes/"+strconv.Itoa(ID))
	respondWithJSON(w, http.StatusOK, "")
}

func (a *App) updateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	if noteExists := a.isNoteExist(id, w); !noteExists {
		return
	}

	var note model.Note
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	updated, err := a.NoteRepository.UpdateNote(id, note)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	w.Header().Add("Location", APIBasePath+"/notes/"+strconv.Itoa(id))
	respondWithJSON(w, http.StatusOK, updated)
}

func (a *App) deleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid note ID")
		return
	}

	if noteExists := a.isNoteExist(id, w); !noteExists {
		return
	}

	removed, err := a.NoteRepository.DeleteNote(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(w, http.StatusOK, removed)

}

func (a *App) isNoteExist(id int, w http.ResponseWriter) bool {
	// TODO: replace by isNoteExists
	_, err := a.NoteRepository.GetNote(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Note with ID: '"+strconv.Itoa(id)+"' Not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return false
	}
	return true
}

func respondWithError(w http.ResponseWriter, statusCode int, errorMsg string) {
	respondWithJSON(w, statusCode, map[string]string{"error": errorMsg})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
