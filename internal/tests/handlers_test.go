package tests

import (
	"bytes"
	"complaint_service/internal/api/handlers"
	"complaint_service/internal/entity"
	"complaint_service/internal/mocks"
	"complaint_service/internal/models"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestFindUsers_Success(t *testing.T) {
	mockProcessor := new(mocks.MockComplaintsProcessor)
	handler := handlers.CreateComplaintsHandler(mockProcessor)

	app := fiber.New()
	app.Get("/user/:id", handler.FindUsers)

	userUUID := uuid.NewV4()
	expectedUser := entity.Users{UserUUID: userUUID, UserName: "John Doe"}

	mockProcessor.On("FindUsers", userUUID.String()).Return(expectedUser, nil)

	req := httptest.NewRequest("GET", "/user/"+userUUID.String(), nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockProcessor.AssertExpectations(t)
}

func TestUpdateComplaintStatus_Success(t *testing.T) {
	mockProcessor := new(mocks.MockComplaintsProcessor)
	handler := handlers.CreateComplaintsHandler(mockProcessor)

	app := fiber.New()
	app.Put("/reports/:id", handler.UpdateComplaintStatus)

	id := "some-complaint-id"
	requestBody := models.Request{Status: "resolved", AdminComment: "Issue resolved"}

	mockProcessor.On("UpdateComplaintStatus", id, requestBody.Status, requestBody.AdminComment).Return(time.Now(), nil)

	reqBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("PUT", "/reports/some-complaint-id", bytes.NewBuffer(reqBody))
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	mockProcessor.AssertExpectations(t)
}
