package handlers

import (
	"complaint_service/internal/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

const (
	successfulReg = "успешная регистрация"
	badRequest    = "incorrect request"
	serverError   = "ошибка сервера"
)

func (h *ComplaintsHandler) signUp(c *fiber.Ctx) {
	var input models.UserSignUp

	if err := c.BodyParser(&input); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSONP(
			models.ResponseSignUp{
				Id:     0,
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
			models.ResponseSignUp{
				Id:     0,
				Status: fmt.Sprintf("%v: %v", serverError, err),
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		models.ResponseSignUp{
			Id:     id,
			Status: successfulReg,
		})
	if err != nil {
		log.Println(err)
	}
}

func (h *ComplaintsHandler) signIn(c *fiber.Ctx) {
	var input models.UserSignUp

	if err := c.BodyParser(&input); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSONP(
			models.ResponseSignIn{
				Token:  "",
				Status: badRequest,
			})
		if err != nil {
			log.Println(err)
		}
		return
	}
	token, err := h.complaintsProcessor.Authorization.GetToken(input.UserName, input.Password)
	if err != nil {
		err = c.Status(fiber.StatusInternalServerError).JSONP(
			models.ResponseSignIn{
				Token:  "",
				Status: fmt.Sprintf("%v: %v", serverError, err),
			})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = c.Status(fiber.StatusOK).JSONP(
		models.ResponseSignIn{
			Token:  token,
			Status: successfulReg,
		})
	if err != nil {
		log.Println(err)
	}
}
