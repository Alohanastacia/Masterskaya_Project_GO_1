GOOSE_DRIVER=postgres
#DB_TYPE=postgres #Тип БД
SSL_MODE=disable #режим подключения к БД
GOOSE_MIGRATION_DIR=./migration #путь к миграциям

#строка подключения к БД
GOOSE_DBSTRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=$(SSL_MODE)

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

#migrate-up применяет все up-миграции goose up
migrate-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) up

#m-last-down откатывает последнюю миграцию
m-last-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) down

#m-status проверяет статус миграции
m-status:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir $(GOOSE_MIGRATION_DIR) status
