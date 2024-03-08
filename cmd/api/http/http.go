package http

import (
	"go-keep/cmd/api"
	"go-keep/cmd/api/http/note"
	"go-keep/internal/config"
	"net/http"
)

func Start(conf *config.Configuration, pkg api.Packager) error {

	note.NewNoteRoute(conf, pkg)
	address := conf.Server.HTTP.Address
	http.ListenAndServe(address, nil)

	return nil
}
