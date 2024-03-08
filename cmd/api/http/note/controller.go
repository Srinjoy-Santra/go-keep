package note

import (
	"encoding/json"
	"errors"
	"go-keep/cmd/api"
	"go-keep/pkg/note"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type NoteService struct {
	pkg api.Packager
}

func NewNoteService(pkg api.Packager) *NoteService {
	return &NoteService{pkg}
}

func (n *NoteService) create(w *http.ResponseWriter, r *http.Request) {

	log.Println("Create request")

	var note note.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	notePkg := n.pkg.NewNotePkg()
	err := notePkg.Create(&note)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusNotImplemented)
	}

	(*w).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*w).Encode(note); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (n *NoteService) get(w *http.ResponseWriter, r *http.Request) {
	log.Println("Get request")

	rId, err := GetId(r.RequestURI)
	(*w).Header().Set("Content-Type", "application/json")
	if err != nil {
		if err := json.NewEncoder(*w).Encode(n); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		id, err := uuid.Parse(rId)
		if err != nil {
			http.Error(*w, err.Error(), http.StatusBadRequest)
		}

		log.Print(id)
		notePkg := n.pkg.NewNotePkg()
		notes, err := notePkg.Get("")
		if err != nil {
			http.Error(*w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(*w).Encode(notes); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetId(uri string) (string, error) {
	subpaths := strings.Split(uri, "/notes/")
	if len(subpaths) > 1 {
		return subpaths[1], nil
	}
	return "", errors.New("no id found")
}
