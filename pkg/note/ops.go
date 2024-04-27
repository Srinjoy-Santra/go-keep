package note

import (
	"go-keep/internal/config"
	"go-keep/pkg/session"
)

type NotePkg struct {
	config *config.Configuration
	opr    operator
	Ss     *session.SessionStore[session.Session]
}

func NewNotePkg(conf *config.Configuration, operator operator, store *session.SessionStore[session.Session]) *NotePkg {
	return &NotePkg{config: conf, opr: operator, Ss: store}
}

func (pkg *NotePkg) Create(n *Note) error {
	err := pkg.opr.Insert(n)
	if err != nil {
		return err
	}
	return nil
}

func (pkg *NotePkg) Get(query, userId string) ([]*Note, error) {
	notes, err := pkg.opr.Get(query, userId)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (pkg *NotePkg) GetOne(id, userId string) (*Note, error) {
	note, err := pkg.opr.GetOne(id, userId)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (pkg *NotePkg) GetAll(userId string) ([]*Note, error) {
	notes, err := pkg.opr.GetAll(userId)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (pkg *NotePkg) Update(n *Note) error {
	err := pkg.opr.Update(n)
	if err != nil {
		return err
	}
	return nil
}

func (pkg *NotePkg) Remove(id, userId string) error {
	err := pkg.opr.Delete(id, userId)
	if err != nil {
		return err
	}
	return nil
}
