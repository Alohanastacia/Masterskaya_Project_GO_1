package entity

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Users struct {
	ID          uint   `db:"id" json:"id"`
	UserUUID    uint   `db:"user_uuid" json:"user_uuid"`
	UserName    string `db:"user_name" json:"user_name"`
	Password    string `db:"password" json:"password"`
	Email       string `db:"email" json:"email"`
	Phone       string `db:"phone" json:"phone"`
	Blacklisted bool   `db:"blacklisted" json:"blacklisted"`
	Role        Role   `db:"role" json:"role"`
}
