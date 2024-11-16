package main

import (
	"github.com/garrickedd/ReLibca/src/server/api"
	"github.com/garrickedd/ReLibca/src/server/api/data/cache"
	"github.com/garrickedd/ReLibca/src/server/config"
	"github.com/garrickedd/ReLibca/src/server/infrastructure/persistence/database"
	"github.com/garrickedd/ReLibca/src/server/infrastructure/persistence/migration"
	"github.com/garrickedd/ReLibca/src/server/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up_1()

	api.InitServer(cfg)

	// db.DbConnection()
	// db.Migrate()
}
