package db

import (
	"errors"
	"fmt"
	"go-keep/internal/config"

	"github.com/google/uuid"
)

func NewDummy(conf *config.Configuration) *Dummy {
	return &Dummy{
		connection: "Put some conf string",
		notes:      []Note{},
	}
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

func (d *Dummy) GetOne(id string) (Note, error) {

	var noteOne Note
	dId, err := uuid.Parse(id)
	if err != nil {
		return noteOne, fmt.Errorf("invalid id %s", id)
	}
	for _, note := range d.notes {
		if note.ID == dId {
			noteOne = note
			break
		}
	}

	return noteOne, nil
}

func (d *Dummy) Get(query string) ([]Note, error) {
	return d.notes, nil
}

func (d *Dummy) GetAll() ([]Note, error) {
	return d.notes, nil
}

func (d *Dummy) Update(note *Note) error {

	for i, lNote := range d.notes {
		if lNote.ID == note.ID {
			d.notes[i] = *note
			return nil

		}
	}
	return errors.New("id not found")
}

func (d *Dummy) Delete(id string) error {

	dId, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid id %s", id)
	}

	for i, note := range d.notes {
		if note.ID == dId {
			d.notes = append(d.notes[:i], d.notes[i+1:]...)
			break
		}
	}

	return errors.New("id not found")
}
