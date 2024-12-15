package handlers

import (
	"complaint_service/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type ComplaintsProcessor interface {
	ComplaintsListAdmin(UserUUID string) (repository.Users, error)
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

// func (h *ComplaintsHandler) ChangeAdminRole(c *fiber.Ctx) error {
func (h *ComplaintsHandler) ComplaintsListAdmin(c *fiber.Ctx) error {
	UserUUID := c.Params("id")
	res, err := h.complaintsProcessor.ComplaintsListAdmin(UserUUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UserUUID is not found"})
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

//var exist bool

//emptyCheck := `SELECT EXISTS(SELECT 1 FROM users WHERE user_uuid = $1)`
//err := repository.Users

//query := `UPDATE users SET role = 'ADMIN' WHERE user_uuid = $1`
//_, err := repository.Users.Exec(query, userID)
//if err != nil {
//	return c.Status(500).SendString("Error updating role")
//}
//return c.JSON(fiber.Map{
//	"status": "роль успешно обновлена",
//})
//}
