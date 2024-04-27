package note

import (
	"encoding/json"
	"errors"
	"go-keep/cmd/api"
	"go-keep/pkg/note"
	"go-keep/pkg/session"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type NoteService struct {
	pkg *note.NotePkg
	ss  *session.SessionStore[session.Session]
}

func NewNoteService(pkg api.Packager) *NoteService {
	notePkg := pkg.NewNotePkg()
	ss := notePkg.Ss
	return &NoteService{notePkg, ss}
}

func (n *NoteService) create(w http.ResponseWriter, r *http.Request) {

	var note note.Note
	userId, err := n.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	note.UserId = userId

	err = n.pkg.Create(&note)
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

func (n *NoteService) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := n.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	filter := r.URL.Query().Get("q")
	if filter != "" {
		notes, err := n.pkg.Get(filter, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}

		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	filter = r.PathValue("id")
	id, err := validateUUId(filter)
	if err != nil {
		notes, err := n.pkg.GetAll(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
		}

		if err := json.NewEncoder(w).Encode(notes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Print(id)
		notes, err := n.pkg.GetOne(filter, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotImplemented)
			return
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

	userId, err := n.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	note.UserId = userId

	id, err := validateUUId(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print(id)
	note.ID = id

	err = n.pkg.Update(&note)
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
		return
	}

	log.Print(id)
	userId, err := n.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	err = n.pkg.Remove(id.String(), userId)
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

func (n *NoteService) getUserId(r *http.Request) (string, error) {
	session := n.ss.GetSessionFromCtx(r)
	sub, ok := session.Profile["sub"]
	if ok {
		userId, ok := sub.(string)
		if ok {
			return userId, nil
		}
	}

	return "", errors.New("user not authenticated")
}
