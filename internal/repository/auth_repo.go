package repository

import (
	"complaint_service/internal/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.Users) (int, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

// Функция NewAuthPostgres является конструктором структуры AuthPostgres. Принимает на вход переменную типа sqlx.DB и возвращает AuthPostgres.
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

/*
Функция CreateUser отправляет INSERT запрос в базу данных для создания пользователя. Принимает на вход структуру User, возвращает переменные id типа int и err типа error
*/
func (r *AuthPostgres) CreateUser(user entity.Users) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (user_uuid,user_name,password,role) values($1,$2,$3,$4) RETURNING id")
	row := r.db.QueryRow(query, user.UserUUID, user.UserName, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
