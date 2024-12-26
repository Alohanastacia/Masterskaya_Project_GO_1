package tests

import (
	"complaint_service/internal/entity"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockComplaintsProcessor struct {
	mock.Mock
}

func (m *MockComplaintsProcessor) FindUsers(UserUUID string) (entity.Users, error) {
	args := m.Called(UserUUID)
	return args.Get(0).(entity.Users), args.Error(1)
}

func (m *MockComplaintsProcessor) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
	args := m.Called(id, status, adminComment)
	return args.Get(0).(time.Time), args.Error(1)
}
