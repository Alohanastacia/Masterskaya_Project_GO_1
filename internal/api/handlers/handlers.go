package handlers

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/processors"

	fiber2 "github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber"
)

type ComplaintsProcessor interface {
	FindUsers(UserUUID string) (entity.Users, error)
	//имплиментируются методы из processors
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
