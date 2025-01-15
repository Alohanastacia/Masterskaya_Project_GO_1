package repository

import (
	"complaint_service/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable   = "users"
	reportsTable = "reports"
)

// NewPostgresDB создаёт подключение к базе данных. Возвращает sqlx.DB и error
func NewPostgresDB() (*sqlx.DB, error) {
	configs, err := config.LoadEnv()
	if err != nil {
		fmt.Println(err)
	}

	dbConnectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBDbname, configs.DBPassword, "disable")

	fmt.Println(dbConnectString)
	db, err := sqlx.Open("postgres", dbConnectString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
