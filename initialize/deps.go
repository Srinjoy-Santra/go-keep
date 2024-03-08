package initialize

import (
	"go-keep/internal/config"
	"go-keep/internal/db"
	"go-keep/pkg/note"
	note_repo "go-keep/pkg/note/repository"
)

type PkgDeps struct {
	conf       *config.Configuration
	dbInstance *db.DBInstance
}

func NewPkgDeps(conf *config.Configuration, dbInstance *db.DBInstance) *PkgDeps {
	pkgDeps := &PkgDeps{conf, dbInstance}
	return pkgDeps
}

func (p *PkgDeps) NewNotePkg() *note.NotePkg {
	noteRepo := note_repo.NewNoteRepo(p.conf, p.dbInstance)
	return note.NewNotePkg(p.conf, noteRepo)
}
