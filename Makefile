# Include .env file
include .env
export $(shell sed 's/=.*//' .env)

# Commands
run:
	@go run cmd/server/main.go
	
db-status:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) status

up: 
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) up

down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) down

delete-all:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) reset

reset:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) reset
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_CONNECTION) goose -dir=$(MIGRATION_PATH) up

seed:
	@go run migrations/seed/main.go