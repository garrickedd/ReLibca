package main

import (
	"github.com/garickedd/ReLibca/src/api"
	"github.com/garickedd/ReLibca/src/api/config"
	"github.com/garickedd/ReLibca/src/infrastructure/persistence/database"
	"github.com/garickedd/ReLibca/src/infrastructure/persistence/migration"
	"github.com/garickedd/ReLibca/src/shared/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up_1()

	api.InitServer(cfg)
}
