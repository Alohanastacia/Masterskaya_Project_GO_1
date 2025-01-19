-- +goose Up
-- +goose StatementBegin
CREATE TYPE priority AS ENUM ('high','medium','low');
CREATE TYPE stage AS ENUM ('new','inprogress','done', 'canceled');

ALTER TABLE reports 
ALTER COLUMN priority TYPE priority USING CASE
    WHEN priority = 1 THEN 'high'::priority
    WHEN priority = 2 THEN 'medium'::priority
    WHEN priority = 3 THEN 'low'::priority
    ELSE NULL
END,
ALTER COLUMN stage TYPE stage USING CASE
    WHEN stage = 1 THEN 'new'::stage
    WHEN stage = 2 THEN 'inprogress'::stage
    WHEN stage = 3 THEN 'done'::stage
    WHEN stage = 4 THEN 'canceled'::stage
    ELSE NULL
END;

ALTER TABLE reports
ALTER COLUMN stage SET DEFAULT 'new'::stage,
ALTER COLUMN uuid TYPE UUID USING uuid::UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS priority;
DROP TYPE IF EXISTS stage;

ALTER TABLE reports 
ALTER COLUMN priority TYPE TEXT,
ALTER COLUMN stage TYPE TEXT;
-- +goose StatementEnd

