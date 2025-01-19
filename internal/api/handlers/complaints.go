package handlers

import (
	"complaint_service/internal/models"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber"
)

const (
	successful        = "успешная регистрация"
	badRequestReport  = "ошибка запроса"
	serverErrorReport = "ошибка сервера"
)

func (h *ComplaintsHandler) CreateReport(c *fiber.Ctx) {
	var input models.Reports
	token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")

	if token == "" {
		log.Println(token)
		c.Status(fiber.StatusBadRequest).JSONP(
			models.ResponseSignUp{
				Id:     0,
				Status: badRequestReport,
			})
	}

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		err = c.Status(fiber.StatusBadRequest).JSONP(
			models.ResponseSignUp{
				Id:     0,
				Status: badRequestReport,
			})
		return
	}

	id, err := h.complaintsProcessor.ComplaintsManager.CreateReport(input, token)

	if err != nil {
		log.Println(err)
		err = c.Status(fiber.StatusInternalServerError).JSONP(
			models.ResponseSignUp{
				Id:     0,
				Status: fmt.Sprintf("%v: %v", serverErrorReport, err),
			})
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		models.ResponseSignUp{
			Id:     id,
			Status: successful,
		})
	if err != nil {
		log.Println(err)
	}
}
