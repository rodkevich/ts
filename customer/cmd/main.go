package main

import (
	"log"
	"os"

	"github.com/rodkevich/ts/customer/config"
	"github.com/rodkevich/ts/customer/pkg/postgres"
)

func main() {

	// configs:
	configPath := config.GetConfigPath(os.Getenv("config"))
	conf, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("unable to parse cfg path %v", err)

	}

	// pg:
	pgxConn, err := postgres.NewPgxConn(conf)
	if err != nil {
		log.Fatal("", err.Error())
	}
	defer pgxConn.Close()

}
