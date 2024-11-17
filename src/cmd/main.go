package main

import (
	"log"

	"github.com/garickedd/ReLibca/src/api"
	"github.com/garickedd/ReLibca/src/api/config"
	database "github.com/garickedd/ReLibca/src/infrastructure/persistence"
)

func main() {
	cfg := config.GetConfig()

	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		log.Fatal()
	}

	api.InitServer()
}
