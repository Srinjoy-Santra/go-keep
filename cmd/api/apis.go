package api

import "go-keep/pkg/note"

type Packager interface {
	NewNotePkg() *note.NotePkg
}
