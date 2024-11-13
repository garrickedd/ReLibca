package main

import (
	"github.com/garrickedd/ReLibca/src/server/api"
	"github.com/garrickedd/ReLibca/src/server/infrastructure/db"
)

func main() {
	api.InitServer()

	db.DbConnection()
	db.Migrate()
}
