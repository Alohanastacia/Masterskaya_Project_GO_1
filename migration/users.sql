CREATE TYPE role AS ENUM ('USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS users(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_uuid uuid GENERATED ALWAYS AS IDENTITY,
    username uuid,
    email TEXT,
    phone INT,
    role role,
    password VARCHAR(32)
);
CREATE INDEX IF NOT EXISTS idx_id ON users (id);

