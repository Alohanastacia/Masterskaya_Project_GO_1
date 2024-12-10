package handlers

import (
	"complaint_service/internal/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ComplaintsProcessor interface {
	ComplaintsListAdmin(userId string) (err error)
	//имплиментируются методы из processors
}

type ComplaintsHandler struct {
	complaintsProcessor ComplaintsProcessor
}

func CreateComplaintsHandler(complaintsProcessor ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики
// Get registers a route for GET methods that requests a representation
// of the specified resource. Requests using GET should only retrieve data.

func (h *ComplaintsHandler) ChangeAdminRole(c *fiber.Ctx) error {
	userID := c.Params("user_uuid")
	query := `UPDATE users SET role = 'ADMIN' WHERE user_uuid = $1`
	_, err := repository.Users{}.Exec(query, userID)
	if err != nil {
		return c.Status(500).SendString("Error updating role")
	}
	return c.JSON(fiber.Map{
		"status": "роль успешно обновлена",
	})
}

func (h *ComplaintsHandler) ComplaintsListAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	user, exists := users[id]
	if !exists {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.JSON(user)

}
