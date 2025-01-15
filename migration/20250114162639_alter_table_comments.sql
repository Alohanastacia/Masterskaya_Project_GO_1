-- +goose Up
-- +goose StatementBegin
ALTER TABLE comments
ADD COLUMN report_id INT,
ADD CONSTRAINT fk_report_id FOREIGN KEY (report_id) REFERENCES reports(Id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP COLUMN report_id;
-- +goose StatementEnd
