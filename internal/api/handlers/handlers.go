package handlers

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/logger"
	"complaint_service/internal/model"
	"complaint_service/internal/processors"

	"log/slog"

	fiber2 "github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber"
)

type ComplaintsProcessor interface {
	FindUsers(UserUUID string) (entity.Users, error)
	//имплиментируются методы из processors
	CreateComplaints(c model.CreateComplaint) (int64, error)
}

type ComplaintsHandler struct {
	complaintsProcessor *processors.ComplaintsProcessor
}

func CreateComplaintsHandler(complaintsProcessor ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики
// Get registers a route for GET methods that requests a representation
// of the specified resource. Requests using GET should only retrieve data.

func (h *ComplaintsHandler) FindUsers(c *fiber2.Ctx) error {
	UserUUID := c.Params("id")
	res, err := h.complaintsProcessor.FindUsers(UserUUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UserUUID is not found"})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// Функция InitRoutes инициализирует роуты. Принимает на вход переменную типа fiber.App
func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
}

func (h *ComplaintsHandler) CreateComplaints(c *fiber2.Ctx) error {

	var request model.CreateComplaint
	log := logger.Log

	const op = "handlers.CreateComplaints"

	log.With(
		slog.String("op", op),
	)

	if err := c.BodyParser(&request); err != nil {
		log.Info("error Bad request")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad request"})
	}

	id, err := h.complaintsProcessor.CreateComplaints(request)

	if err != nil {
		log.Info("Error saving to database")

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error saving to database"})
	}

	return c.Status(fiber.StatusOK).JSON(id)
}
