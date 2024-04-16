package initialize

import (
	"go-keep/internal"
	"go-keep/internal/config"
	"go-keep/internal/db"
	"go-keep/pkg/note"
	note_repo "go-keep/pkg/note/repository"
	"go-keep/pkg/session"
	"time"
)

type PkgDeps struct {
	conf         *config.Configuration
	dbInstance   *db.DBInstance
	sessionStore *session.SessionStore[session.Session]
	auth         *internal.Authenticator
}

func NewPkgDeps(conf *config.Configuration, dbInstance *db.DBInstance, auth *internal.Authenticator) *PkgDeps {

	var ss session.SessionStore[session.Session]
	ss.InitStore("SessionID", time.Duration(time.Hour*24*7)) // 1 week

	pkgDeps := &PkgDeps{conf, dbInstance, &ss, auth}
	return pkgDeps
}

func (p *PkgDeps) NewNotePkg() *note.NotePkg {
	noteRepo := note_repo.NewNoteRepo(p.conf, p.dbInstance)
	return note.NewNotePkg(p.conf, noteRepo)
}

func (p *PkgDeps) NewUserPkg() *session.UserPkg {
	return session.NewUserPkg(p.conf, p.sessionStore, p.auth)
}
