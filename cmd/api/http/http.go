package http

import (
	"go-keep/cmd/api"
	"go-keep/cmd/api/http/note"
	"go-keep/internal/config"
	"log"
	"net/http"
)

func Start(conf *config.Configuration, pkg api.Packager) error {

	router := http.NewServeMux()

	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	note.NewNoteRoute(conf, pkg, router)
	server := http.Server{
		Addr:    conf.Server.HTTP.Address,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())

	return nil
}
