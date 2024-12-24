-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  category_report (
    report_id INT,
    category_id INT,
    FOREIGN KEY (report_id) REFERENCES reports (id)
    FOREIGN KEY (category_id) REFERENCES category (id)
)
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category_reports;
-- +goose StatementEnd
