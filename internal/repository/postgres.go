package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sqlx.DB, error) {
	dbConnectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost", "5432", "postgres", "complaint_service", "qwerty", "disable")

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
