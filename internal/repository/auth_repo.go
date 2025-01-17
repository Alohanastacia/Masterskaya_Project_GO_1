package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.UserSignUp) (int, error)
	GetUser(username, password string) (entity.Users, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres является конструктором структуры AuthPostgres. Принимает на вход переменную типа sqlx.DB и возвращает AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

/*
CreateUser отправляет INSERT запрос в базу данных для создания пользователя. Принимает на вход структуру User,
возвращает переменные id типа int и err типа error
*/
func (r *AuthPostgres) CreateUser(userModel models.UserSignUp) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	user := entity.Users{
		UserName: userModel.UserName,
		Password: userModel.Password,
		UserUUID: userModel.UserUUID,
		Role:     entity.Role(models.User),
	}
	query := fmt.Sprintf("INSERT INTO users (user_uuid,username,password,role) values($1,$2,$3,$4) RETURNING id")
	row := tx.QueryRow(query, user.UserUUID, user.UserName, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return id, nil
}

/*
GetUser отправляет SELECT запрос в базу данных для получения данных пользователя. Принимает на вход username и password,
возвращает структуру User и ошибку типа error
*/
func (r *AuthPostgres) GetUser(username, password string) (entity.Users, error) {
	var user entity.Users
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
