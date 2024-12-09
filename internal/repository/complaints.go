package repository

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"time"
)

type Complaint struct {
	ID          uint      `db:"id"`
	UUID        string    `db:"uuid"`
	UserUUID    string    `db:"user_uuid"`
	Description string    `db:"description"`
	Priority    string    `db:"priority"`
	Stage       string    `db:"stage"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Users struct {
	ID       uint   `db:"id"`
	UserId   uint   `db:"user_id"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Status   string `db:"status"`
}

type ComplaintsRepository struct {
	db *sqlx.DB
}

func CreateComplaintsRepository(db *sqlx.DB) *ComplaintsRepository {
	return &ComplaintsRepository{db: db}
}

func (rep *ComplaintsRepository) ComplaintsListAdmin(userId string) ([]User, error) {
	var users []User

	if err := rep.db.Get(&users, `SELECT * FROM users WHERE user_id=? AND role = 'ADMIN'`, userId); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return users, nil
}

// Ниже будут методы ComplaintsRepository, которые делают запросы в БД и отдают результат
