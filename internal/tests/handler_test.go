package test

//
//import (
//	"bytes"
//	"encoding/json"
//	"net/http/httptest"
//	"testing"
//	"time"
//
//	"complaint_service/internal/api/handlers"
//	"complaint_service/internal/entity"
//	"complaint_service/internal/mocks"
//	"github.com/gofiber/fiber/v2"
//	"github.com/satori/go.uuid"
//	"github.com/stretchr/testify/assert"
//)
//
//func TestFindUsers(t *testing.T) {
//	app := fiber.New()
//	mockProcessor := new(mocks.MockComplaintsProcessor)
//
//	handler := handlers.CreateComplaintsHandler(mockProcessor)
//
//	app.Get("/api/v1/user/:id", handler.FindUsers)
//
//	userUUID := uuid.NewV4()
//	mockProcessor.On("FindUsers", userUUID.String()).Return(entity.Users{UserUUID: userUUID, UserName: "test-user"}, nil)
//
//	req := httptest.NewRequest("GET", "/api/v1/user/"+userUUID.String(), nil)
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, 200, resp.StatusCode)
//
//	var user entity.Users
//	json.NewDecoder(resp.Body).Decode(&user)
//	assert.Equal(t, "test-user", user.UserName)
//
//	mockProcessor.AssertExpectations(t)
//}
//
//func TestUpdateComplaintStatus(t *testing.T) {
//	app := fiber.New()
//	mockProcessor := new(mocks.MockComplaintsProcessor)
//
//	handler := handlers.CreateComplaintsHandler(mockProcessor)
//
//	app.Put("/api/v1/reports/:id", handler.UpdateComplaintStatus)
//
//	reqBody := []byte(`{"status": "DONE", "adminComment": "Approved"}`)
//	req := httptest.NewRequest("PUT", "/api/v1/reports/test-id", bytes.NewBuffer(reqBody))
//	resp, _ := app.Test(req)
//
//	mockProcessor.On("UpdateComplaintStatus", "test-id", "DONE", "Approved").Return(time.Now(), nil)
//
//	assert.Equal(t, 200, resp.StatusCode)
//
//	mockProcessor.AssertExpectations(t)
//}
//
//func TestDeleteComment(t *testing.T) {
//	app := fiber.New()
//	mockProcessor := new(mocks.MockComplaintsProcessor)
//
//	handler := handlers.CreateComplaintsHandler(mockProcessor)
//
//	app.Delete("/reports/:id/comments/:commentId", handler.DeleteComment)
//
//	mockProcessor.On("DeleteComment", "test-id", "comment-id").Return(nil)
//
//	req := httptest.NewRequest("DELETE", "/reports/test-id/comments/comment-id", nil)
//	resp, _ := app.Test(req)
//
//	assert.Equal(t, 204, resp.StatusCode)
//
//	mockProcessor.AssertExpectations(t)
//}
//
//func TestUpdateComplaintPriority(t *testing.T) {
//	app := fiber.New()
//	mockProcessor := new(mocks.MockComplaintsProcessor)
//
//	handler := handlers.CreateComplaintsHandler(mockProcessor)
//
//	app.Put("/api/v1/reports/:id/priority", handler.UpdateComplaintPriority)
//
//	reqBody := []byte(`{"priority": "Высокий"}`)
//	req := httptest.NewRequest("PUT", "/api/v1/reports/test-id/priority", bytes.NewBuffer(reqBody))
//	resp, _ := app.Test(req)
//
//	mockProcessor.On("UpdateComplaintPriority", "test-id", "Высокий").Return(time.Now(), nil)
//
//	assert.Equal(t, 200, resp.StatusCode)
//
//	mockProcessor.AssertExpectations(t)
//}
