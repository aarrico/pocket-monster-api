include .env

GOOSE_DBSTRING='postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)'
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR='db/migrations'

build:
	go build -o out/pokemon-api main.go

migrate:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

rollback-one:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

reset-db:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down-to 0

seed: build
	out/pokemon-api seed
