package repo

import (
	"go-keep/internal/config"
	"go-keep/internal/db"
	pn "go-keep/pkg/note"
)

type dbInstancer interface {
	GetDB() db.Dber
}

type NoteRepo struct {
	config *config.Configuration
	db     db.Dber
}

func NewNoteRepo(conf *config.Configuration, dbInstances dbInstancer) *NoteRepo {
	return &NoteRepo{conf, dbInstances.GetDB()}
}

func (nr *NoteRepo) Insert(n *pn.Note) error {
	note := db.Note{
		Title:    n.Title,
		Content:  n.Content,
		UserName: n.UserName,
	}
	err := nr.db.Insert(&note)
	if err != nil {
		return err
	}
	n.ID = note.ID
	return nil
}

func (nr *NoteRepo) Get(query, userId string) ([]*pn.Note, error) {
	dbn, err := nr.db.Get(query, userId)
	if err != nil {
		return nil, err
	}
	return bindToNotes(dbn), nil
}

func (nr *NoteRepo) GetOne(id, userId string) (*pn.Note, error) {
	dbn, err := nr.db.GetOne(id, userId)
	if err != nil {
		return nil, err
	}
	return bindToNote(dbn), nil
}

func (nr *NoteRepo) GetAll(userId string) ([]*pn.Note, error) {
	dbn, err := nr.db.GetAll(userId)
	if err != nil {
		return nil, err
	}
	return bindToNotes(dbn), nil
}

func (nr *NoteRepo) Update(n *pn.Note) error {
	dn := db.Note{
		ID:       n.ID,
		Title:    n.Title,
		Content:  n.Content,
		UserName: n.UserName,
	}
	return nr.db.Update(&dn)
}

func (nr *NoteRepo) Delete(id string) error {
	return nr.db.Delete(id)
}

func bindToNotes(dbNotes []db.Note) []*pn.Note {
	notes := []*pn.Note{}
	for _, v := range dbNotes {
		notes = append(notes, bindToNote(v))
	}
	return notes
}

func bindToNote(dbNote db.Note) *pn.Note {
	return &pn.Note{
		ID:       dbNote.ID,
		Title:    dbNote.Title,
		Content:  dbNote.Content,
		UserName: dbNote.UserName,
	}
}
