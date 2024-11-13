include .env

GOOSE_DBSTRING='postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)'
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR='db/migrations'
SEED_ALL?=false

build:
	go build -o out/pokemon-api main.go

start-db:
	docker-compose up -d db

check-db:
	@pg_isready -d $(GOOSE_DBSTRING) -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) || $(MAKE) start-db

migrate: check-db
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

rollback-one: check-db
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

reset-db: check-db
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down-to 0

seed: check-db
	go build -o out/seed-db cmd/seed_db/main.go
	if [ "$(SEED_ALL)" = "true" ]; then \
   		$(MAKE) reset-db; \
   		$(MAKE) migrate; \
 	fi
	out/seed-db $(SEED_ALL)
