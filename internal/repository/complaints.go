package repository

import "github.com/jmoiron/sqlx"

type ComplaintsRepository struct {
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db: db}
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат
