package db

import (
	"database/sql"
	"errors"
	"fmt"
	"go-keep/internal/config"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func NewPostgres(conf *config.Configuration) (*Postgres, error) {

	connection := conf.Database.Relational.Connection
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Postgres{connection: connection, Db: db}, nil
}

type Postgres struct {
	connection string
	Db         *sql.DB
}

// Insert implements Dber.
func (p *Postgres) Insert(note *Note) error {
	id, err := uuid.NewV7()
	if err == nil {
		query := `INSERT INTO note (id, title, content) VALUES ($1, $2, $3) RETURNING id`

		var pk uuid.UUID
		err := p.Db.QueryRow(query, id, note.Title, note.Content).Scan(&pk)
		if err != nil {
			return err
		}
		note.ID = pk
		return nil
	} else {
		return err
	}
}

// GetOne implements Dber.
func (p *Postgres) GetOne(id string) (Note, error) {

	var note Note
	pk, err := uuid.Parse(id)
	if err != nil {
		return note, fmt.Errorf("invalid id %s", id)
	}

	query := `SELECT title, content From note WHERE id = $1`
	var title, content string
	err = p.Db.QueryRow(query, pk).Scan(&title, &content)
	if err != nil {
		return note, err
	}

	return Note{
		ID:      pk,
		Title:   title,
		Content: content,
	}, nil
}

// Get implements Dber.
func (p *Postgres) Get(query string) ([]Note, error) {
	panic("unimplemented")
}

// GetAll implements Dber.
func (p *Postgres) GetAll() ([]Note, error) {
	notes := []Note{}

	query := `SELECT id, title, content From note`
	rows, err := p.Db.Query(query)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	var id uuid.UUID
	var title, content string
	for rows.Next() {
		err := rows.Scan(&id, &title, &content)
		if err != nil {
			return notes, err
		}
		notes = append(notes, Note{
			ID:      id,
			Title:   title,
			Content: content,
		})
	}

	return notes, nil

}

// Update implements Dber.
func (p *Postgres) Update(note *Note) error {
	query := `UPDATE note SET title=$2, content=$3 WHERE id=$1`

	res, err := p.Db.Exec(query, note.ID, note.Title, note.Content)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	} else if count != 1 {
		return errors.New("1 record was not updated")
	}

	return nil
}

// Delete implements Dber.
func (p *Postgres) Delete(id string) error {

	pk, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid id %s", id)
	}

	query := `DELETE FROM note WHERE id=$1`

	res, err := p.Db.Exec(query, pk)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	} else if count != 1 {
		errors.New("1 record was not deleted")
	}

	return nil
}
