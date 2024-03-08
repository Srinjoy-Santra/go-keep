package note

import (
	"go-keep/internal/config"
)

type NotePkg struct {
	config *config.Configuration
	opr    operator
}

func NewNotePkg(conf *config.Configuration, operator operator) *NotePkg {
	return &NotePkg{config: conf, opr: operator}
}

func (pkg *NotePkg) Create(n *Note) error {
	/* 	_ := Note{
		Title:   n.Title,
		Content: n.Content,
	} */
	err := pkg.opr.Insert(n)
	if err != nil {
		return err
	}
	return nil
}

func (pkg *NotePkg) Get(query string) ([]*Note, error) {
	notes, err := pkg.opr.Get(query)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
