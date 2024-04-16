package initialize

import (
	"fmt"
	"go-keep/cmd/api/http"
	"go-keep/internal"
	"go-keep/internal/config"
	"go-keep/internal/db"
	"log"
	"sync"
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
	dbInstances, err := db.NewInitializedInstances(conf)
	if err != nil {
		return err
	}

	auth, err := internal.NewAuth(conf)
	if err != nil {
		return err
	}

	StartServers(conf, dbInstances, auth)
	return nil
}

func StartServers(conf *config.Configuration, dbInstance *db.DBInstance, auth *internal.Authenticator) {

	pkg := NewPkgDeps(conf, dbInstance, auth)
	var wg sync.WaitGroup
	wg.Add(1)
	go startHTTP(&wg, conf, pkg)
	//wg.Add(1)
	//go startGRPC(&wg, conf, pkg, errCh)
	wg.Wait()
}

func startHTTP(wg *sync.WaitGroup, conf *config.Configuration, pkg *PkgDeps) {
	defer wg.Done()
	err := http.Start(conf, pkg)
	if err != nil {
		log.Println(err)
	}
}

func startGRPC(wg *sync.WaitGroup, conf *config.Configuration, pkg *PkgDeps) {
	defer wg.Done()
}
