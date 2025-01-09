
# Указываем базовый образ
FROM golang:1.23.3-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o main ./cmd/main.go

# Финальный образ
FROM debian:bullseye-slim

# Копируем собранный бинарный файл из builder-образа
COPY --from=builder /app/main .

# Копируем файлы конфигурации
COPY ./.env ./.env
COPY ./config ./config

# Указываем порт, который будет использовать приложение
EXPOSE 8080

# Указываем команду запуска
CMD ["./main"]
