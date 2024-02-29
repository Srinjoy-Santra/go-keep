package init

import (
	"fmt"
	"go-keep/internal/config"
)

func Run() error {
	env, err := Getenv()
	if err != nil {
		return err
	}

	fmt.Println("Environment : ", env)

	// Initialize configs
	conf, err := config.NewConfig(env)
	if err != nil {
		return err
	}

	// Initialize db
	dbInstances, err := db.NewInstances(conf)
	if err != nil {

	}
}
