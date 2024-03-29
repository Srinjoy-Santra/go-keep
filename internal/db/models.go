package db

import "github.com/google/uuid"

type Note struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type Dber interface {
	inserter
	getter
	updater
	deleter
}

type inserter interface {
	Insert(*Note) error
}
type getter interface {
	GetOne(id string) (Note, error)
	Get(query string) ([]Note, error)
	GetAll() ([]Note, error)
}

type updater interface {
	Update(*Note) error
}

type deleter interface {
	Delete(string) error
}
