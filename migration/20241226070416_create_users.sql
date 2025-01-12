-- +goose Up
-- +goose StatementBegin
CREATE TYPE role AS ENUM ('USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS users(
       id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       user_uuid UUID NOT NULL UNIQUE,
       username UUID NOT NULL UNIQUE,
       email TEXT,
       phone INT,
       role role DEFAULT 'USER',
       password VARCHAR(32)
    );
CREATE INDEX IF NOT EXISTS idx_id ON users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TYPE IF EXISTS role;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
ALTER TABLE user(CREATE role SuperAdmin WITH role 'Cap' WITH LOGIN PASSWORD 'pwd' CREATEROLE);
-- +goose StatementEnd
-- + goose Down
-- +goose StatementBegin
ALTER TABLE user DELETE FROM users WHERE username = 'Cap';;
-- +goose StatementEnd

/*
 Талица users.
 1. Добавлена вторая миграция. Создание супер администратора с логином и паролем.
 Определены его права выдавать роль администратора.
 */