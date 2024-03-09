package note

import (
	"go-keep/cmd/api"
	"go-keep/internal/config"
	"log"
	"net/http"
	"strings"
)

func NewNoteRoute(conf *config.Configuration, pkg api.Packager) {
	service := NewNoteService(pkg)

	noteHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received at", r.URL.Path)
		path := strings.Trim(r.URL.Path, "/")
		paths := strings.Split(path, "/")
		resource := paths[0]

		switch resource {
		case "ping":
			w.Write([]byte("pong"))
		case "notes":
			switch r.Method {
			case "POST":
				service.create(&w, r)
				return
			case "GET":
				service.get(&w, r)
				return
			case "DELETE":
				service.remove(&w, r)
				return
			case "PUT":
				service.update(&w, r)
				return
			}
		default:
			http.Error(w, "Route not matched", http.StatusUnauthorized)
		}
	}

	http.HandleFunc("/", noteHandler)
}
