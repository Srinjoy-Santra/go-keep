package db

import (
	"go-keep/internal/config"
)

const (
	InMemory   = "InMemory"
	Relational = "Relational"
)

type DBInstance struct {
	conf    config.Configuration
	dummyDB *Dummy
	pgDB    *Postgres
}

func NewInitializedInstances(conf *config.Configuration) (*DBInstance, error) {
	dbInstance := &DBInstance{conf: *conf}
	err := dbInstance.initialize()
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

func (i *DBInstance) initialize() error {
	if i.conf.Database.Source == InMemory {
		i.dummyDB = NewDummy(&i.conf)
	} else {
		i.pgDB = NewPostgres(&i.conf)
	}
	return nil
}

func (i *DBInstance) GetDB() Dber {
	if i.conf.Database.Source == InMemory {
		return i.dummyDB
	} else {
		return i.pgDB
	}
}
