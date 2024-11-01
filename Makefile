# Migration
GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING ?= "postgres://postgres:admin@localhost:5432/relibcadb?sslmode=disable"
GOOSE_MIGRATION_DIR ?= "./src/api/migrations"

up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

.PHONY: up down reset