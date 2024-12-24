package model

import "time"

type CreateComplaint struct {
	Priority    string    `json:"priority"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Created_at  time.Time `json:"created_at"`
}
