package db

import (
	"errors"
	"fmt"
	"go-keep/internal/config"

	"github.com/google/uuid"
)

func NewDummy(conf *config.Configuration) (*Dummy, error) {
	return &Dummy{
		connection: "Put some conf string",
		notes:      []Note{},
	}, nil
}

type Dummy struct {
	connection string
	notes      []Note
}

func (d *Dummy) Insert(note *Note) error {
	id, err := uuid.NewV7()
	if err == nil {
		note.ID = id
		d.notes = append(d.notes, *note)
		return nil
	} else {
		return err
	}
}

func (d *Dummy) GetOne(id, userId string) (Note, error) {

	var noteOne Note
	dId, err := uuid.Parse(id)
	if err != nil {
		return noteOne, fmt.Errorf("invalid id %s", id)
	}
	for _, note := range d.notes {
		if note.ID == dId && note.UserId == userId {
			noteOne = note
			break
		}
	}

	return noteOne, nil
}

func (d *Dummy) Get(query, userId string) ([]Note, error) {
	return d.notes, nil
}

func (d *Dummy) GetAll(userId string) ([]Note, error) {
	if userId == "admin" {
		return d.notes, nil
	}

	notes := []Note{}
	for _, note := range d.notes {
		if note.UserId == userId {
			notes = append(notes, note)
		}
	}

	return notes, nil

}

func (d *Dummy) Update(note *Note) error {

	for i, lNote := range d.notes {
		if lNote.ID == note.ID && lNote.UserId == note.UserId {
			d.notes[i] = *note
			return nil

		}
	}
	return errors.New("id not found")
}

func (d *Dummy) Delete(id, userId string) error {

	dId, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid id %s", id)
	}

	for i, note := range d.notes {
		if note.ID == dId && note.UserId == userId {
			d.notes = append(d.notes[:i], d.notes[i+1:]...)
			break
		}
	}

	return errors.New("id not found")
}
