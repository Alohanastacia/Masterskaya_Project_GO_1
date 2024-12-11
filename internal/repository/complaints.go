package repository

import (
	"complaint_service/internal/model"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ComplaintsRepository struct {
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db: db}
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат

func (rep *ComplaintsRepository) GetComplaint(uuid string) (model.GetComplaint, error) {
	var getComplaint model.GetComplaint
	query := `SELECT UUID, stage, priority, description, created_at FROM reports WHERE UUID = :uuid`
	rows := rep.db.QueryRow(query, sql.Named("uuid", uuid))
	if err := rows.Scan(&getComplaint.UUID, &getComplaint.Stage, &getComplaint.Priority, &getComplaint.Description, &getComplaint.Created_at); err != nil {
		return getComplaint, fmt.Errorf("UUID is not found %w", err)
	}
	return getComplaint, nil
}
