package entity

import uuid "github.com/satori/go.uuid"

type Role string

const (
	SuperAdmin Role = "SUPER_ADMIN"
	User       Role = "USER"
	Admin      Role = "ADMIN"
)

type Users struct {
	ID       uint      `db:"id" json:"id"`
	UserUUID uuid.UUID `db:"user_UUID" json:"user_UUID"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Email    string    `db:"email" json:"email"`
	Phone    string    `db:"phone" json:"phone"`
	Role     Role      `db:"role" json:"role"`
}

type SuperUser struct {
	ID        uint      `db:"id" json:"id"`
	UserUUID  uuid.UUID `db:"user_UUID" json:"user_UUID"`
	AdminName string    `db:"admin_name" json:"admin_name"`
	Password  string    `db:"password" json:"password"`
	Role      Role      `db:"role" json:"role"`
}
