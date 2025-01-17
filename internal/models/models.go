package models

import (
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

type RequestStatustics struct {
	Period    string `json:"period"`     //Available values : day, week, month
	StartDate string `json:"start_date"` //Начальная дата для фильтрации (YYYY-MM-DD)
	EndDate   string `json:"end_date"`   //Конечная дата для фильтрации (YYYY-MM-DD)
	Status    string `json:"status"`     //Available values : NEW, IN_PROGRESS, DONE, CANCELLED
	Category  string `json:"category"`   //Фильтр по категории
	Limit     int    `json:"limit"`      //Максимальное количество записей на странице (по умолчанию 10)
	Offset    int    `json:"offset"`     //Смещение для пагинации (по умолчанию 0)
}

type ResponseStatistics struct {
	TotalReports      int            `json:"total_reports"`
	ReportsByStatus   map[string]int `json:"reports_by_status"`
	ReportsByCategory map[string]int `json:"reports_by_category"`
}
