CREATE TYPE report_stage AS ENUM ('IN_PROGRESS', 'CANCELLED', 'DONE', 'MODERATION');

CREATE TABLE IF NOT EXISTS reports_history (
    id SERIAL PRIMARY KEY,
    report_id UUID REFERENCES reports(id) ON DELETE CASCADE,
    old_stage report_stage NOT NULL,
    new_stage report_stage NOT NULL,
    admin_comment TEXT,
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX IF NOT EXISTS idx_report_id ON reports_history (report_id);
