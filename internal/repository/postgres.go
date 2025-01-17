package repository

import (
	"complaint_service/internal/config"
	"fmt"
	"log"
	"time"

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
		fmt.Println("не удалось загрузить переменные окружения для связи с БД:", err)
	}

	dbConnectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBDbname, configs.DBPassword, "disable")
	fmt.Println(dbConnectString)

	db, err := sqlx.Open("postgres", dbConnectString)
	if err != nil {
		return nil, err
	}

	tStartPing := time.Now()
	err = db.Ping()
	tEndPing := time.Now()
	if err != nil {
		log.Println("Пинг не выполнен")
		return nil, err
	}
	log.Println("Пинг выполнен за:", tEndPing.Sub(tStartPing))

	return db, nil
}
