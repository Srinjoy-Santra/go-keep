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
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
		return
	}

	notePkg := n.pkg.NewNotePkg()
	err := notePkg.Create(&note)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
	}

	(*w).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*w).Encode(note); err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
		return
	}

	(*w).WriteHeader(http.StatusCreated)
}

func (n *NoteService) get(w *http.ResponseWriter, r *http.Request) {
	log.Println("Get request")
	(*w).Header().Set("Content-Type", "application/json")
	notePkg := n.pkg.NewNotePkg()

	filter := r.URL.Query().Get("q")
	if filter != "" {
		notes, err := notePkg.Get(filter)
		if err != nil {
			http.Error(*w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(*w).Encode(notes); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	filter, err := getId(r.RequestURI)
	if err != nil {

		notes, err := notePkg.GetAll()
		if err != nil {
			http.Error(*w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(*w).Encode(notes); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		id, err := uuid.Parse(filter)
		if err != nil {
			http.Error(*w, err.Error(), http.StatusBadRequest)
		}

		log.Print(id)
		notes, err := notePkg.Get(filter)
		if err != nil {
			http.Error(*w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(*w).Encode(notes); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (n *NoteService) remove(w *http.ResponseWriter, r *http.Request) {
	log.Println("Remove request")

	rId, err := getId(r.RequestURI)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
	}

	id, err := uuid.Parse(rId)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
	}
	log.Print(id)

	notePkg := n.pkg.NewNotePkg()
	err = notePkg.Remove(rId)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).WriteHeader(http.StatusNoContent)
}

func (n *NoteService) update(w *http.ResponseWriter, r *http.Request) {

	log.Println("Update request")

	var note note.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
		return
	}

	rId, err := getId(r.RequestURI)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(rId)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
	}
	log.Print(id)
	note.ID = id

	notePkg := n.pkg.NewNotePkg()
	err = notePkg.Update(&note)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
	}

	(*w).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*w).Encode(note); err != nil {
		http.Error(*w, err.Error(), http.StatusNoContent)
		return
	}

	(*w).WriteHeader(http.StatusCreated)
}

func getId(uri string) (string, error) {
	subpaths := strings.Split(uri, "/notes/")
	if len(subpaths) > 1 {
		return subpaths[1], nil
	}
	return "", errors.New("no id found")
}
