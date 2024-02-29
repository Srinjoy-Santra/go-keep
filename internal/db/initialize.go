package db

import (
	"go-keep/internal/config"
)

type DBInstance struct {
	conf config.Configuration
	myDB *MyDB
}

func NewInitializedInstances(conf *config.Configuration) (*DBInstance[], error) {
	dbInstance := &DBInstance{conf: conf}
	err:= dbInstance.initialize()
	if err !=nil {
		return nil
	}

}
