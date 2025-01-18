package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Role string

const (
	User  Role = "USER"
	Admin Role = "ADMIN"
)

type UserSessions struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type ResponseSignUp struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

type ResponseSignIn struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

type UserSignUp struct {
	UserUUID uuid.UUID `json:"user_UUID"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
}

type Request struct {
	Status       string `json:"status"`
	AdminComment string `json:"admin_comment"`
}

type UpdateStatusResponse struct {
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

var (
	ErrCommentNotFound   = errors.New("comment not found")
	ErrComplaintNotFound = errors.New("complaint not found")
)
