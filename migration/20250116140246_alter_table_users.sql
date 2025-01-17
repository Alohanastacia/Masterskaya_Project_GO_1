-- +goose Up
ALTER TABLE users 
ALTER COLUMN username TYPE TEXT,
ALTER COLUMN password TYPE TEXT;


-- +goose Down
ALTER TABLE users 
ALTER COLUMN username TYPE UUID USING NULLIF(username, '')::UUID,
ALTER COLUMN password TYPE VARCHAR(32);
