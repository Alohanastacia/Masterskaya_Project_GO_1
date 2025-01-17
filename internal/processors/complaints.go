package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"complaint_service/internal/repository"
	"github.com/satori/go.uuid"
	"time"
)

type ComplaintsRepository interface {
	FindUsers(UserUUID uuid.UUID) (*entity.Users, error)
	UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error)
	DeleteComment(complaintID string, commentID string) error
	UpdateComplaintPriority(id string, priority string) (time.Time, error)
}

type ComplaintsProcessor struct {
	Authorization
	complaintsRepository *repository.ComplaintsRepository
}

// CreateComplaintsProcessor является конструктором структуры ComplaintsProcessor.
func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization:        NewAuthService(complaintsRepository.Authorization),
		complaintsRepository: complaintsRepository,
	}
}

func (p *ComplaintsProcessor) FindUsers(UserUUID uuid.UUID) (entity.Users, error) {
	user, err := p.complaintsRepository.FindUsers(UserUUID)
	if err != nil {
		return entity.Users{}, err
	}
	return *user, nil
}

func (p *ComplaintsProcessor) UpdateComplaintStatus(id string, status string, adminComment string) (time.Time, error) {
	return p.complaintsRepository.UpdateComplaintStatus(id, status, adminComment)
}
func (p *ComplaintsProcessor) DeleteComment(complaintID string, commentID string) error {
	return p.complaintsRepository.DeleteComment(complaintID, commentID)
}

func (p *ComplaintsProcessor) UpdateComplaintPriority(id string, priority string) (time.Time, error) {
	return p.complaintsRepository.UpdateComplaintPriority(id, priority)
}

// Implement Authorization methods
func (p *ComplaintsProcessor) CreateUser(user models.UserSignUp) (int, error) {
	return p.Authorization.CreateUser(user)
}

func (p *ComplaintsProcessor) GetToken(username, password string) (string, error) {
	return p.Authorization.GetToken(username, password)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику и вызываются из хендлеров
