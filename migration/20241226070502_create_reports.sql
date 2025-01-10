-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reports (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	uuid UUID,
	user_uuid UUID,
	description VARCHAR(255),
	priority INT,
	stage INT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);
CREATE INDEX IF NOT EXISTS idx_created_at ON reports (created_at);
CREATE INDEX IF NOT EXISTS idx_id ON reports (id);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS reports;

/*
reports - таблица жалоб
--------------------------------------------------
	priority		приоритет пока INT, а нужно ('high','medium','low')
	stage			стадия обработки пока INT, а нужно ('new','inprogress','done', 'canceled')

Оставляю поля по которым есть вопросы.
Варианты решения: 
1. С помощью ENUM 
2. С помощью дополнительных таблиц и внешних ключей FOREIN KEY (...) REFERENCES ... 	
*/