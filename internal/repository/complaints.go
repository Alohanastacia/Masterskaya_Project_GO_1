package repository

import (
	"complaint_service/internal/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ComplaintsRepository struct {
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db: db}
}

func (rep *ComplaintsRepository) FindUsers(UserUUID string) ([]*entity.Users, error) {

	var user entity.Users
	const query = `SELECT user_uuid, username,email, role, phone,blacklisted 
					FROM users 
					WHERE user_uuid = ?
					ORDER BY user_uuid 
					OFFSET ? ROWS FETCH NEXT 10 ROWS ONLY;`

	if UserUUID == "" {
		return nil, fmt.Errorf("user_uuid is required")
	}
	rows := rep.db.QueryRow(query, UserUUID)
	err := rows.Scan(
		&user.UserUUID,
		&user.UserName,
		&user.Email,
		&user.Role,
		&user.Phone,
		&user.Blacklisted,
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
