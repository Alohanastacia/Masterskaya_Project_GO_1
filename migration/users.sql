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
