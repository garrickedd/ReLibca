package main

import (
	"log"

	"github.com/garrickedd/ReLibca/src/server/api"
	"github.com/garrickedd/ReLibca/src/server/api/data/cache"
	"github.com/garrickedd/ReLibca/src/server/config"
	database "github.com/garrickedd/ReLibca/src/server/infrastructure/persistence"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		log.Fatal(err)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)

	// db.DbConnection()
	// db.Migrate()
}
