package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type Priority string

const (
	High   Priority = "high"
	Medium Priority = "medium"
	Low    Priority = "low"
)

type Stage string

const (
	New        Stage = "new"
	Inprogress Stage = "inprogress"
	Done       Stage = "done"
	Canceled   Stage = "canceled"
)

type Users struct {
	ID       uint      `db:"id" json:"id"`
	UserUUID uuid.UUID `db:"user_uuid" json:"user_UUID"`
	UserName string    `db:"username" json:"user_name"`
	Password string    `db:"password" json:"password"`
	Email    string    `db:"email" json:"email"`
	Phone    int       `db:"phone" json:"phone"`
	Role     Role      `db:"role" json:"role"`
}

type Reports struct {
	Uuid        uuid.UUID `json:"UUID"`
	User_uuid   uuid.UUID `json:"user_UUID"`
	Description string    `json:"description"`
	Proirity    Priority  `json:"priority"`
	Stage       Stage     `json:"stage"`
}
