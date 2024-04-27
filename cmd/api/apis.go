package api

import (
	"go-keep/pkg/note"
	"go-keep/pkg/session"
)

type Packager interface {
	NewNotePkg() *note.NotePkg
	NewUserPkg() *session.UserPkg
}
