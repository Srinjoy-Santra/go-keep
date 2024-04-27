package user

import (
	"go-keep/cmd/api"
	"go-keep/internal/config"
	"go-keep/pkg/session"
	"net/http"
)

func NewUserRoute(conf *config.Configuration, pkg api.Packager, router *http.ServeMux) *session.SessionStore[session.Session] {
	service := NewUserService(pkg)

	router.HandleFunc("GET /user", service.user)
	router.HandleFunc("GET /login", service.login)
	router.HandleFunc("GET /callback", service.callback)
	router.HandleFunc("GET /logout", service.logout)

	return service.Ss
}
