package note

import (
	"encoding/json"
	"go-keep/cmd/api"
	"go-keep/pkg/note"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type NoteService struct {
	pkg api.Packager
}

func NewNoteService(pkg api.Packager) *NoteService {
	return &NoteService{pkg}
}

func (n *NoteService) create(w http.ResponseWriter, r *http.Request) {

	var note note.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	notePkg := n.pkg.NewNotePkg()
	err := notePkg.Create(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (n *NoteService) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	notePkg := n.pkg.NewNotePkg()

	filter := r.URL.Query().Get("q")
	if filter != "" {
		notes, err := notePkg.Get(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	filter = r.PathValue("id")
	id, err := validateUUId(filter)
	if err != nil {
		notes, err := notePkg.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Print(id)
		notes, err := notePkg.GetOne(filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
		}

		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (n *NoteService) update(w http.ResponseWriter, r *http.Request) {

	var note note.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	id, err := validateUUId(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(id)
	note.ID = id

	notePkg := n.pkg.NewNotePkg()
	err = notePkg.Update(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (n *NoteService) remove(w http.ResponseWriter, r *http.Request) {

	id, err := validateUUId(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	log.Print(id)

	notePkg := n.pkg.NewNotePkg()
	err = notePkg.Remove(id.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func validateUUId(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}
	return uid, nil
}
