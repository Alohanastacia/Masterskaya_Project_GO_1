package handlers

import (
	"complaint_service/internal/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	successfulReg = "успешная регистрация"
	badRequest    = "неправильные данные запроса"
	serverError   = "ошибка сервера"
)

func (h *ComplaintsHandler) signUp(c *fiber.Ctx) error {
	var input models.UserSignUp

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseSignUp{
			Id:     0,
			Status: badRequest,
		})
	}

	id, err := h.complaintsProcessor.CreateUser(input)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseSignUp{
			Id:     0,
			Status: fmt.Sprintf("%v: %v", serverError, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseSignUp{
		Id:     id,
		Status: successfulReg,
	})
}

func (h *ComplaintsHandler) signIn(c *fiber.Ctx) error {
	var input models.UserSignUp

	if err := c.BodyParser(&input); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ResponseSignIn{
			Token:  "",
			Status: badRequest,
		})
	}

	token, err := h.complaintsProcessor.GetToken(input.UserName, input.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.ResponseSignIn{
			Token:  "",
			Status: fmt.Sprintf("%v: %v", serverError, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.ResponseSignIn{
		Token:  token,
		Status: successfulReg,
	})
}
