package note

import (
	"go-keep/cmd/api"
	"go-keep/internal/config"
	"net/http"
)

func NewNoteRoute(conf *config.Configuration, pkg api.Packager, router *http.ServeMux) {
	service := NewNoteService(pkg)

	router.HandleFunc("GET /notes/", service.get)
	router.HandleFunc("GET /notes/{id}", service.get)

	router.HandleFunc("POST /notes", service.create)
	router.HandleFunc("PUT /notes/{id}", service.update)

	router.HandleFunc("DELETE /notes/{id}", service.remove)
}
