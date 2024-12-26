package handlers

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/processors"
	"net/http"
	"time"

	fiber2 "github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber"
)

const (
	StatusUpdatedMessage = "успешно обновлено"
	ErrorInvalidRequest  = "Invalid request body"
	ErrorUpdatingStatus  = "Error updating complaint status"
	ErrorCommentNotFound = "Комментарий не найден"
	ErrorDeletingComment = "Ошибка при удалении комментария"
	//ErrorForbidden       = "Доступ запрещен"
)

type ComplaintsProcessor interface {
	FindUsers(UserUUID string) (entity.Users, error)
	UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error)
	DeleteComment(complaintID string, commentID string) error

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

func (h *ComplaintsHandler) DeleteComment(c *fiber2.Ctx) error {
	id := c.Params("id")
	commentId := c.Params("commentId")

	// Проверка прав администратора
	//if !isAdmin(c) { // функция проверки прав администратора (пока не разобрался с этим)
	//	return c.Status(http.StatusForbidden).JSON(fiber.Map{"error": ErrorForbidden})
	//}

	err := h.complaintsProcessor.DeleteComment(id, commentId)
	if err != nil {
		if err == entity.ErrCommentNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": ErrorCommentNotFound})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": ErrorDeletingComment})
	}

	c.SendStatus(http.StatusNoContent) // 204 No Content
	return nil
}

func (h *ComplaintsHandler) UpdateComplaintStatus(c *fiber2.Ctx) error {
	id := c.Params("id")

	var request entity.Request

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": ErrorInvalidRequest})
	}

	updatedAt, err := h.complaintsProcessor.UpdateComplaintStatus(id, request.Status, request.AdminComment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": ErrorUpdatingStatus})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     StatusUpdatedMessage,
		"updated_at": updatedAt.Format(time.RFC3339),
	})
}

// Функция InitRoutes инициализирует роуты. Принимает на вход переменную типа fiber.App
func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
}
