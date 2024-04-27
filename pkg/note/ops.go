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

func (pkg *NotePkg) Remove(id, userName string) error {
	err := pkg.opr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
