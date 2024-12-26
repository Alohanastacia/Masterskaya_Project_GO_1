package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"time"
)

type ComplaintsRepository interface {
	FindUsers(UserUUID string) ([]*entity.Users, error)
	UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error)
	DeleteComment(complaintID string, commentID string) error

	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	Authorization
}

func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization: NewAuthService(complaintsRepository.Authorization),
	}
}

func (p *ComplaintsProcessor) FindUsers(UserUUID string) (entity.Users, error) {
	return p.FindUsers(UserUUID)
}

func (p *ComplaintsProcessor) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
	return p.UpdateComplaintStatus(id, status, adminComment)
}

func (p *ComplaintsProcessor) DeleteComment(complaintID string, commentID string) error {
	return p.DeleteComment(complaintID, commentID)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
