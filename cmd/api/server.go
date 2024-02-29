package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Note struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type Server struct {
	Notes []Note
}

func NewServer() *Server {
	s := &Server{
		Notes: []Note{},
	}

	return s
}

func (s *Server) CreateNote(w *http.ResponseWriter, r *http.Request) {

	log.Println("Create request")

	var n Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.NewV7()
	if err == nil {
		n.ID = id
	} else {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.Notes = append(s.Notes, n)

	(*w).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*w).Encode(s.Notes); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *Server) GetNotes(w *http.ResponseWriter, r *http.Request) {
	log.Println("Get request")

	rId, err := GetId(r.RequestURI)
	(*w).Header().Set("Content-Type", "application/json")
	if err != nil {
		if err := json.NewEncoder(*w).Encode(s.Notes); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		id, err := uuid.Parse(rId)
		if err != nil {
			http.Error(*w, err.Error(), http.StatusBadRequest)
		}

		log.Print(id)
		var noteOne Note
		for _, note := range s.Notes {
			if note.ID == id {
				noteOne = note
				break
			}
		}

		if err := json.NewEncoder(*w).Encode(noteOne); err != nil {
			http.Error(*w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) RemoveNote(w *http.ResponseWriter, r *http.Request) {
	log.Println("Remove request")

	rId, err := GetId(r.RequestURI)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
	}

	id, err := uuid.Parse(rId)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
	}
	log.Print(id)

	for i, note := range s.Notes {
		if note.ID == id {
			s.Notes = append(s.Notes[:i], s.Notes[i+1:]...)
			break
		}
	}

	(*w).Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(*w).Encode(s.Notes); err != nil {
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetId(uri string) (string, error) {
	subpaths := strings.Split(uri, "/notes/")
	if len(subpaths) > 1 {
		return subpaths[1], nil
	}
	return "", errors.New("no id found")
}
