package main

import (
	"github.com/garrickedd/ReLibca/src/server/api"
	"github.com/garrickedd/ReLibca/src/server/api/data/cache"
	"github.com/garrickedd/ReLibca/src/server/config"
	"github.com/garrickedd/ReLibca/src/server/infrastructure/db"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)

	db.DbConnection()
	db.Migrate()
}
