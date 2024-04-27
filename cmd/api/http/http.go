package http

import (
	"go-keep/cmd/api"
	"go-keep/cmd/api/http/middleware"
	"go-keep/cmd/api/http/note"
	"go-keep/cmd/api/http/user"
	"go-keep/internal/config"
	"log"
	"net/http"
)

func Start(conf *config.Configuration, pkg api.Packager) error {

	router := http.NewServeMux()

	router.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	ss := user.NewUserRoute(conf, pkg, router)

	authRouter := http.NewServeMux()
	note.NewNoteRoute(conf, pkg, authRouter)
	router.Handle("/", ss.LoadSession(authRouter))

	server := http.Server{
		Addr:    conf.Server.HTTP.Address,
		Handler: middleware.Logging(router),
	}
	log.Fatal(server.ListenAndServe())

	return nil
}
