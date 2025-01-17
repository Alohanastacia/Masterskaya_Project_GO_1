package mocks

//
//import (
//	"complaint_service/internal/entity"
//	"complaint_service/internal/models"
//	uuid "github.com/satori/go.uuid"
//	"github.com/stretchr/testify/mock"
//	"time"
//)
//
//type MockComplaintsProcessor struct {
//	mock.Mock
//}
//
//func (m *MockComplaintsProcessor) FindUsers(UserUUID uuid.UUID) (entity.Users, error) {
//	args := m.Called(UserUUID)
//	return args.Get(0).(entity.Users), args.Error(1)
//}
//
//func (m *MockComplaintsProcessor) CreateUser(user models.UserSignUp) (int, error) {
//	args := m.Called(user)
//	return args.Int(0), args.Error(1)
//}
//
//func (m *MockComplaintsProcessor) GetToken(username, password string) (string, error) {
//	args := m.Called(username, password)
//	return args.String(0), args.Error(1)
//}
//
//func (m *MockComplaintsProcessor) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
//	args := m.Called(id, status, adminComment)
//	return args.Get(0).(time.Time), args.Error(1)
//}
//
//func (m *MockComplaintsProcessor) DeleteComment(complaintID string, commentID string) error {
//	args := m.Called(complaintID, commentID)
//	return args.Error(0)
//}
//
//func (m *MockComplaintsProcessor) UpdateComplaintPriority(id string, priority string) (time.Time, error) {
//	args := m.Called(id, priority)
//	return args.Get(0).(time.Time), args.Error(1)
//}
