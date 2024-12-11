package handlers

import (
	"complaint_service/internal/model"

	"github.com/gofiber/fiber/v2"
)

type ComplaintsProcessor interface {
	//имплиментируются методы из processors
	GetComplaint(uuid string) (model.GetComplaint, error)
}

type ComplaintsHandler struct {
	complaintsProcessor ComplaintsProcessor
}

func CreateComplaintsHandler(complaintsProcessor ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики
func (h *ComplaintsHandler) GetComplaint(c *fiber.Ctx) error {
	uuid := c.Params("id")
	res, err := h.complaintsProcessor.GetComplaint(uuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UUID is not found"})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
