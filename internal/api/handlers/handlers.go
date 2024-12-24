package handlers

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/logger"
	"complaint_service/internal/processors"
	"log/slog"

	"github.com/gofiber/fiber"
)

type ComplaintsProcessor interface {
	//имплиментируются методы из processors
	CreateComplaints(c entity.CreateComplaint) (int64, error)
}

type ComplaintsHandler struct {
	complaintsProcessor *processors.ComplaintsProcessor
}

// Функция CreateComplaintsHandler является конструктором структуры ComplaintsHandler. Принимает на вход переменную типа processors.ComplaintsProcessor и возвращает ComplaintsHandler.
func CreateComplaintsHandler(complaintsProcessor *processors.ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor: complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики

// Функция InitRoutes инициализирует роуты. Принимает на вход переменную типа fiber.App
func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
}

func (h *ComplaintsHandler) CreateComplaints(c *fiber.Ctx) error {

	var request entity.CreateComplaint
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
