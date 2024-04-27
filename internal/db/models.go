package db

import "github.com/google/uuid"

type Note struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	UserId  string    `json:"userId"`
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
	GetOne(id, userId string) (Note, error)
	Get(query, userId string) ([]Note, error)
	GetAll(userId string) ([]Note, error)
}

type updater interface {
	Update(*Note) error
}

type deleter interface {
	Delete(id, userId string) error
}
