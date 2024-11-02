# Migration
GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING ?= "postgres://postgres:admin@localhost:5432/relibcadb?sslmode=disable"
GOOSE_MIGRATION_DIR ?= "./src/api/migrations"

postgres:
	docker run --name docker-db-1 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres:alpine

migrateup:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
migratedown:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: postgres migrateup migratedown reset