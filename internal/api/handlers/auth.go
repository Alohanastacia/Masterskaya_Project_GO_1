package handlers

import (
	"complaint_service/internal/entity"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

const (
	successfulReg = "успешная регистрация"
	badRequest    = "неправильные данные запроса"
	serverError   = "ошибка севера"
)

func (h *ComplaintsHandler) signUp(c *fiber.Ctx) {
	var input entity.User

	type Response struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSONP(
			Response{
				ID:     0,
				Status: badRequest,
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	id, err := h.complaintsProcessor.Authorization.CreateUser(input)

	if err != nil {
		err = c.Status(fiber.StatusInternalServerError).JSONP(
			Response{
				ID:     0,
				Status: fmt.Sprintf("%v: %v", serverError, err),
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		Response{
			ID:     id,
			Status: successfulReg,
		})
	if err != nil {
		log.Println(err)
	}
}
