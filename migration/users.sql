CREATE TYPE role AS ENUM ('USER', 'ADMIN');

CREATE TABLE IF NOT EXISTS user(
       id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
       user_uuid UUID NOT NULL UNIQUE,
       user_name UUID NOT NULL UNIQUE,
       email TEXT,
       phone INT,
       role role DEFAULT 'USER',
       password VARCHAR(32)
    );
CREATE INDEX IF NOT EXISTS idx_id ON user (id);
