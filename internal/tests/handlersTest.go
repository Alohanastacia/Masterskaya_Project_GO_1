package tests

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"complaint_service/internal/api/handlers"
	"complaint_service/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUpdateComplaintStatus_Success(t *testing.T) {
	app := fiber.New()
	mockProcessor := new(mocks.MockComplaintsProcessor)
	handler := handlers.ComplaintsHandler{ComplaintsProcessor: mockProcessor}

	mockProcessor.On("UpdateComplaintStatus", "123", "resolved", "Handled by admin").Return(time.Now(), nil)

	reqBody := `{"status": "resolved", "admin_comment": "Handled by admin"}`
	req := httptest.NewRequest("PUT", "/api/v1/reports/123", bytes.NewBufferString(reqBody))
	res, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	mockProcessor.AssertExpectations(t)
}

func TestUpdateComplaintStatus_InvalidRequest(t *testing.T) {
	app := fiber.New()
	handler := handlers.ComplaintsHandler{}

	reqBody := `{invalid json}`
	req := httptest.NewRequest("PUT", "/api/v1/reports/123", bytes.NewBufferString(reqBody))
	res, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}
