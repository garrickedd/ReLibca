package main

import "github.com/garrickedd/ReLibca/src/server/infrastructure/db"

func main() {
	db.DbConnection()
	db.Migrate()
}