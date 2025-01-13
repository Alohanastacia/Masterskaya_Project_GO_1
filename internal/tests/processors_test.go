package tests

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/mocks"
	"complaint_service/internal/processors"
	"github.com/satori/go.uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindUsers_Success(t *testing.T) {
	mockRepo := new(mocks.MockComplaintsRepository)
	proc := processors.CreateComplaintsProcessor(mockRepo)

	userUUID := uuid.NewV4() // Generate a new UUID
	expectedUser := entity.Users{UserUUID: userUUID}

	mockRepo.On("FindUsers", userUUID).Return([]*entity.Users{&expectedUser}, nil)

	user, err := proc.FindUsers(userUUID.String())

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.UserUUID, user.UserUUID)

	mockRepo.AssertExpectations(t)
}

func TestUpdateComplaintStatus_Success(t *testing.T) {
	mockRepo := new(mocks.MockComplaintsRepository)
	proc := processors.CreateComplaintsProcessor(mockRepo)

	id := "some-complaint-id"
	status := "resolved"
	adminComment := "Issue resolved"

	mockRepo.On("UpdateComplaintStatus", id, status, adminComment).Return(time.Now(), nil)

	updatedAt, err := proc.UpdateComplaintStatus(id, status, adminComment)

	assert.NoError(t, err)
	assert.NotZero(t, updatedAt)

	mockRepo.AssertExpectations(t)
}
