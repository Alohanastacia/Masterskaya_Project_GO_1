package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"time"
)

type ComplaintsRepository interface {
	FindUsers(UserUUID string) (*entity.Users, error)
	UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error)
	DeleteComment(complaintID string, commentID string) error
	UpdateComplaintPriority(id string, priority string) (time.Time, error)
}

type ComplaintsProcessor struct {
	Authorization
	repo ComplaintsRepository
}

// CreateComplaintsProcessor является конструктором структуры ComplaintsProcessor.
func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization: NewAuthService(complaintsRepository.Authorization),
		repo:          complaintsRepository,
	}
}

// FindUsers ищет пользователя по UserUUID и возвращает его данные.
func (p *ComplaintsProcessor) FindUsers(UserUUID string) (entity.Users, error) {
	user, err := p.repo.FindUsers(UserUUID)
	if err != nil {
		return entity.Users{}, err // Возвращаем пустую структуру в случае ошибки
	}
	return *user, nil // Возвращаем найденного пользователя
}

func (p *ComplaintsProcessor) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
	return p.repo.UpdateComplaintStatus(id, status, adminComment)
}
func (p *ComplaintsProcessor) DeleteComment(complaintID string, commentID string) error {
	return p.repo.DeleteComment(complaintID, commentID)
}

func (p *ComplaintsProcessor) UpdateComplaintPriority(id string, priority string) (time.Time, error) {
	return p.repo.UpdateComplaintPriority(id, priority)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику и вызываются из хендлеров
