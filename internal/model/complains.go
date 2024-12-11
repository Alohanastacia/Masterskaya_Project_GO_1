package model

import "time"

type GetComplaint struct {
	UUID        string    `json:"uuid"`
	Stage       string    `json:"stage"`
	Priority    string    `json:"priority"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
}
