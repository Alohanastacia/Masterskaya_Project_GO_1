package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Id        int       `json:"id"`
	User_UUID uuid.UUID `json:"user_UUID"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
}

type CreateComplaint struct {
	Priority    string    `json:"priority"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Created_at  time.Time `json:"created_at"`
}
