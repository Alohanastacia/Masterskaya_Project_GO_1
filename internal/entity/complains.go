package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

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

type Complaint struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UserUUID    uuid.UUID `db:"user_uuid" json:"user_uuid"`
	Description string    `db:"description" json:"description"`
	Priority    string    `db:"priority" json:"priority"`
	Stage       string    `db:"stage" json:"stage"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type ReportsHistory struct {
	ID           uint      `db:"id" json:"id"`
	ReportID     uuid.UUID `db:"report_id" json:"report_id"`
	OldStage     string    `db:"old_stage" json:"old_stage"`
	NewStage     string    `db:"new_stage" json:"new_stage"`
	AdminComment string    `db:"admin_comment" json:"admin_comment"`
	ChangedAt    time.Time `db:"changed_at" json:"changed_at"`
}
