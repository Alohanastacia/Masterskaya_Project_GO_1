package entity

import uuid "github.com/satori/go.uuid"

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Users struct {
	ID       uint      `db:"id" json:"id"`
	UserUUID uuid.UUID `db:"user_UUID" json:"user_UUID"`
	UserName string    `db:"user_name" json:"user_name"`
	Password string    `db:"password" json:"password"`
	Email    string    `db:"email" json:"email"`
	Phone    string    `db:"phone" json:"phone"`
	Role     Role      `db:"role" json:"role"`
}
