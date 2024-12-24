package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/logger"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type ComplaintsRepository struct {
	Authorization
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{
		Authorization: NewAuthPostgres(db),
		db:            db,
	}
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат

func (rep *ComplaintsRepository) CreateComplaints(c entity.CreateComplaint) (int64, error) {
	const op = "internal.api.repository.CreateComplaints"
	log := logger.Log
	log.With(
		slog.String("op", op),
	)

	equre := `"INSERT INTO reports ( Priority, Description, Category, Created_at) VALUES (?, ?, ?, ?)"`
	res, err := rep.db.Exec(equre, c.Priority, c.Description, c.Category, c.Created_at)
	if err != nil {
		log.Info("Error adding to database")
		return 0, fmt.Errorf("Error adding to database")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Info("Error receiving id")
		return 0, fmt.Errorf("Error receiving id")
	}
	return id, err
}
