package routes

import (
	"complaint_service/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func Complaints(app *fiber.App, complaintsHandler *handlers.ComplaintsHandler) {
	// Пример как задавать роуты
	//app.Get("api/v1/ping", complaintsHandler.GetComplaints)
	app.Get("api/v1/user/{id}", complaintsHandler.FindUsers)
}
