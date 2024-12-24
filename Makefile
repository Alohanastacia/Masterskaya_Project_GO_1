SSL_MODE=disable #отсутствует в конфиге, режим подключения
DB_TYPE=postgres #отсутствует в конфиге, тип БД
MIGRATION_DIR=migration #путь к миграциям

#строка подключения к БД
DB_URL=$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=$(SSL_MODE)

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

#migrate-up применяет все up-миграции
migrate-up:
	goose -dir $(MIGRATION_DIR) up $(DB_URL)

#m-last-down откатывает последнюю миграцию
m-last-down:
	goose -dir $(MIGRATION_DIR) down $(DB_URL)

#m-status проверяет статус миграции
m-status:
	goose -dir $(MIGRATION_DIR) status $(DB_URL)