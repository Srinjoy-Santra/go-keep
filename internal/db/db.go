package db

import "github.com/google/uuid"

type Dber interface {
	getter
	inserter
	updater
	deleter
}

type Note struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type getter interface {
	GetOne(id string) Note
	Get(query string) []Note
	GetAll() []Note
}

type inserter interface {
	Insert(Note) error
}

type updater interface {
	Update(Note) error
}

type deleter interface {
	Delete(string) error
}
