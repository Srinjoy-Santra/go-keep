package note

import "github.com/google/uuid"

type Note struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserName string    `json:"userName"`
}

type operator interface {
	getter
	inserter
	updater
	deleter
}

type getter interface {
	GetOne(id, userId string) (*Note, error)
	Get(query, userId string) ([]*Note, error)
	GetAll(userId string) ([]*Note, error)
}

type inserter interface {
	Insert(*Note) error
}

type updater interface {
	Update(*Note) error
}

type deleter interface {
	Delete(string) error
}
