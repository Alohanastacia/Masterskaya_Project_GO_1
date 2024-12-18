package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Users struct {
	ID       uint   `db:"id" json:"id"`
	UserUUID uint   `db:"user_uuid" json:"user_uuid"`
	UserName string `db:"user_name" json:"user_name"`
	Password string `db:"password" json:"password"`
	Role     Role   `db:"role" json:"role"`
}

type ComplaintsRepository struct {
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db: db}
}

func (rep *ComplaintsRepository) ComplaintsListAdmin(user_uuid string) (Users, error) {
	var users Users
	if users.Role != `ADMIN` {
		return users, fmt.Errorf("access errors, insufficient rights")
	}
	query := `SELECT user_uuid FROM users WHERE user_uuid=?`

	if user_uuid == "" {
		return users, fmt.Errorf("user_uuid is required")
	}
	rows := rep.db.QueryRow(query, user_uuid)
	err := rows.Scan(&users.UserUUID)
	if err != nil {
		return users, fmt.Errorf("user_uuid not found")
	}
	return users, nil
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат
