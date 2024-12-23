package repository

import (
	"complaint_service/internal/entity"
	"fmt"

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
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{
		Authorization: NewAuthPostgres(db),
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
