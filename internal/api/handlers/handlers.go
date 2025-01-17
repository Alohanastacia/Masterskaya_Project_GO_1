package handlers

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"complaint_service/internal/processors"
	"github.com/gofiber/fiber/v2"
	"github.com/satori/go.uuid"
	"time"
)

const (
	StatusUpdatedMessage = "успешно обновлено"
	ErrorInvalidRequest  = "Invalid request body"
	ErrorUpdatingStatus  = "Error updating complaint status"
	ErrorCommentNotFound = "Комментарий не найден"
	ErrorDeletingComment = "Ошибка при удалении комментария"
	//ErrorForbidden       = "Доступ запрещен"
	ErrorUpdatingPriority = "ошибка при обновлении приоритета"
	SuccessPriorityUpdate = "приоритет успешно обновлён"
)

type ComplaintsProcessor interface {
	FindUsers(UserUUID uuid.UUID) (entity.Users, error)
	//CreateUser(user models.UserSignUp) (int, error)
	//GetToken(username, password string) (string, error)
	UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error)
	DeleteComment(complaintID string, commentID string) error
	UpdateComplaintPriority(id string, priority string) (time.Time, error)
}

type ComplaintsHandler struct {
	complaintsProcessor *processors.ComplaintsProcessor
}

func CreateComplaintsHandler(complaintsProcessor *processors.ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor: complaintsProcessor}
}

func (h *ComplaintsHandler) FindUsers(c *fiber.Ctx) error {
	userUUIDStr := c.Params("id")
	userUUID, err := uuid.FromString(userUUIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid UUID format")
	}

	res, err := h.complaintsProcessor.FindUsers(userUUID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "UserUUID not found")
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *ComplaintsHandler) UpdateComplaintStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var request models.Request
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, ErrorInvalidRequest)
	}

	updatedAt, err := h.complaintsProcessor.UpdateComplaintStatus(id, request.Status, request.AdminComment)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, ErrorUpdatingStatus)
	}

	return c.Status(fiber.StatusOK).JSON(models.UpdateStatusResponse{
		Status:    StatusUpdatedMessage,
		UpdatedAt: updatedAt.Format(time.RFC3339),
	})
}

func (h *ComplaintsHandler) DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	commentId := c.Params("commentId")

	// Проверка прав администратора
	// if !isAdmin(c) {
	//	return fiber.NewError(fiber.StatusForbidden, ErrorForbidden)
	// }

	err := h.complaintsProcessor.DeleteComment(id, commentId)
	if err != nil {
		if err == models.ErrCommentNotFound {
			return fiber.NewError(fiber.StatusNotFound, ErrorCommentNotFound)
		}
		return fiber.NewError(fiber.StatusInternalServerError, ErrorDeletingComment)
	}

	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}

func (h *ComplaintsHandler) UpdateComplaintPriority(c *fiber.Ctx) error {
	id := c.Params("id")

	// Проверка прав администратора
	//if !isAdmin(c) {
	//	return fiber.NewError(fiber.StatusForbidden, ErrorForbidden
	//}

	var request struct {
		Priority string `json:"priority"`
	}

	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "некорректный запрос")
	}

	if request.Priority != "Высокий" && request.Priority != "Средний" && request.Priority != "Низкий" {
		return fiber.NewError(fiber.StatusBadRequest, "приоритет должен быть одним из: Высокий, Средний, Низкий")
	}

	updatedAt, err := h.complaintsProcessor.UpdateComplaintPriority(id, request.Priority)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, ErrorUpdatingPriority)
	}

	return c.Status(fiber.StatusOK).JSON(models.UpdateStatusResponse{
		Status:    SuccessPriorityUpdate,
		UpdatedAt: updatedAt.Format(time.RFC3339),
	})
}

// InitRoutes инициализирует маршруты
func (h *ComplaintsHandler) InitRoutes(app *fiber.App) {
	app.Post("user/register", h.signUp)
	app.Post("user/login", h.signIn)
	app.Get("/api/v1/user/:id", h.FindUsers)
	app.Put("api/v1/reports/:id", h.UpdateComplaintStatus)
	app.Delete("/reports/:id/comments/:commentId", h.DeleteComment)
	app.Put("/api/v1/reports/:id/priority", h.UpdateComplaintPriority)

}
