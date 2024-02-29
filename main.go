package main

import (
	"go-keep/cmd/api"
	"log"
	"net/http"
	"strings"
)

func main() {

	c := multiplex{
		api.NewServer(),
	}
	http.Handle("/", &c)

	http.ListenAndServe(":8080", nil)
}

type multiplex struct {
	*api.Server
}

// Use mux / httpRouter
func (m *multiplex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			m.Server.CreateNote(&w, r)
			return
		case "GET":
			m.Server.GetNotes(&w, r)
			return
		case "DELETE":
			m.Server.RemoveNote(&w, r)
			return

		}
	default:
		http.Error(w, "Route not matched", http.StatusUnauthorized)
	}
}
