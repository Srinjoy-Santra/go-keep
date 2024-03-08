package db

import (
	"database/sql"
	"go-keep/internal/config"
)

func NewPostgres(conf *config.Configuration) *Postgres {
	return &Postgres{connection: conf.Database.Relational.Connection}
}

type Postgres struct {
	connection string
	Db         *sql.DB
}

// Delete implements Dber.
func (p *Postgres) Delete(string) error {
	panic("unimplemented")
}

// Get implements Dber.
func (p *Postgres) Get(query string) ([]Note, error) {
	panic("unimplemented")
}

// GetAll implements Dber.
func (p *Postgres) GetAll() ([]Note, error) {
	panic("unimplemented")
}

// GetOne implements Dber.
func (p *Postgres) GetOne(id string) (Note, error) {
	panic("unimplemented")
}

// Insert implements Dber.
func (p *Postgres) Insert(*Note) error {
	panic("unimplemented")
}

// Update implements Dber.
func (p *Postgres) Update(Note) error {
	panic("unimplemented")
}
