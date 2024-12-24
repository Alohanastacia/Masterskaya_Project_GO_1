package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/logger"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

const (
	defaultOffset = 0
	defaultLimit  = 10
)

type ComplaintsDB struct {
	db *sqlx.DB
}

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

func (rep *ComplaintsDB) FindUsers(UserUUID string, limit, offset int) ([]*entity.Users, error) {

	var user entity.Users

	if limit <= 0 {
		limit = defaultLimit
	}
	if offset < 0 {
		offset = defaultOffset
	}

	const query = `SELECT user_uuid, username, email, role, phone
					FROM users 
					WHERE user_uuid = ?
					ORDER BY user_uuid
					LIMIT ? OFFSET ?`

	if UserUUID == "" {
		return nil, fmt.Errorf("user_uuid is required")
	}
	rows := rep.db.QueryRow(query, UserUUID, limit, offset)

	err := rows.Scan(
		&user.UserUUID,
		&user.UserName,
		&user.Email,
		&user.Role,
		&user.Phone,
	)
	if user.Role != entity.Admin {
		return nil, fmt.Errorf("access errors, insufficient rights")
	}
	if err != nil {
		return nil, fmt.Errorf("user_uuid not found")
	}

	return []*entity.Users{&user}, nil
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
