# Migration
GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING ?= "postgres://postgres:admin@localhost:5432/relibcadb?sslmode=disable"
GOOSE_MIGRATION_DIR ?= "./db/migration"

migrateup:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
migratedown:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: migrateup migratedown reset