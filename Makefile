GOOSE_DRIVER=postgres #драйвер библиотеки миграций
SSL_MODE=disable #режим подключения к БД
GOOSE_MIGRATION_DIR=./migration #путь к миграциям

#строка подключения к БД
GOOSE_DBSTRING=://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DBNAME)?sslmode=$(SSL_MODE)

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml
	install-uber-mock:	
	go	install	go.uber.org/mock/mockgen@latest	
mock-processors:	
	mockgen	-source=internal/processors/complaints.go	-destination=internal/processors/mocks/mocks.go