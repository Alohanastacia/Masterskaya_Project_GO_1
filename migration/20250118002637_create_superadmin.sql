-- +goose Up
-- +goose StatementBegin
CREATE ROLE SuperAdmin WITH LOGIN PASSWORD 'pwd' CREATEROLE;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
