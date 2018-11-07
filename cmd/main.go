package main

import (
	"flag"
	"gogo/postgres"
	"log"

	"gogo/config"
)

func main() {
	configPath := flag.String("config", "", "path to configuration file")
	flag.Parse()

	var cfg *config.Config
	var err error
	if *configPath == "" {
		cfg, err = config.FromEnv()
	} else {
		cfg, err = config.FromFile(*configPath)
	}
	if err != nil {
		log.Fatalf("could not obtain configuration: %v", err)
	}

	service, err := postgres.New(cfg.Postgres)
	defer service.DB.Close()

}
