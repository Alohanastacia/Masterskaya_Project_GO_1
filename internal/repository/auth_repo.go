package repository

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.Users) (int, error)
}

func (a Authorization) GetUser(username string, s string) (any, error) {
	panic("unimplemented")
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
func (r *AuthPostgres) CreateUser(user entity.Users) (int, error) {
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("repository.CreateUser ошибка создания транзакции: %v", err)
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
	op := "GetUser"
	log.Println("Старт", op)

	log.Printf("Входящий username: %s, password: %s", username, password)
	var user entity.Users
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	log.Println("Сформирована строка подключения:", query)

	err := r.db.Get(&user, query, username, password)
	if err != nil {
		log.Printf("%s: %s", op, err)
	}
	log.Println("Получены данные:", user)

	return user, err
}
