-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  category (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(32),
    created_at DATE NOT NULL DEFAULT CURRENT_DATE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd