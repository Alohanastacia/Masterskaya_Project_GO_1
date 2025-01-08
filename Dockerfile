# FROM устанавливает версию Go в контейнере
FROM golang:1.23 AS builder

# WORKDIR устанавливает рабочий каталог контейнера
WORKDIR /report-service
RUN pwd
RUN ls -la

# COPY оптимизирует сборку контейнера при изменениях в ПО
COPY go.mod go.sum ./
RUN go mod download
RUN pwd
RUN ls -la

# COPY копирует исходный код
COPY . .
RUN pwd
RUN ls -la

#RUN компилирует приложение app в бинарный файл для Linux x86_64 (amd64)
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go
RUN pwd
RUN ls -la
RUN ls -la /report-service/


# FROM создаёт контейнер на основе минималистичной Linux
FROM alpine:latest
RUN pwd
RUN ls -la

# COPY копирует бинарный app из предыдущего шага в текущую рабочую директорию контейнера
COPY --from=builder /report-service/app .
RUN pwd
RUN ls -la

# COPY копирует файлы конфугураций, необходимые для пакета config
COPY --from=builder /report-service/.env ./
COPY --from=builder /report-service/config/local.yaml ./config/local.yaml
RUN pwd
RUN ls -la

# EXPOSE сообщает ОС, какой порт привязали к процессу контейнера
EXPOSE ${APP_PORT}
RUN pwd
RUN ls -la

# Запускаем приложение
CMD ["./app"]
