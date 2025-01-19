package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/models"
	"complaint_service/internal/repository"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

type ComplaintsManager interface {
	FindUsers(UserUUID string) (entity.Users, error)
	CreateReport(input models.Reports, token string) (int, error)
}

type ComplaintsProcessor struct {
	Authorization
	ComplaintsManager
}

type ComplaintsService struct {
	repo repository.ComplaintsManager
}

func NewComplaintsService(repo repository.ComplaintsManager) *ComplaintsService {
	return &ComplaintsService{
		repo: repo}
}

// CreateComplaintsProcessor является конструктором структуры ComplaintsProcessor. Принимает на вход переменную типа sqlx.DB и возвращает ComplaintsProcessor
func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization:     NewAuthService(complaintsRepository.Authorization),
		ComplaintsManager: NewComplaintsService(complaintsRepository.ComplaintsManager),
	}
}

func (p *ComplaintsService) FindUsers(UserUUID string) (entity.Users, error) {
	return p.FindUsers(UserUUID)
}

func (s *ComplaintsService) CreateReport(input models.Reports, token string) (int, error) {
	if len(input.Description) == 0 || len(input.Proirity) == 0 {
		return 0, fmt.Errorf("Строка жалобы и приоритета не могут быть пустыми")
	}

	user, err := ParseJWT(token)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	report := entity.Reports{
		Uuid:        uuid.NewV4(),
		Description: input.Description,
		Proirity:    entity.Priority(input.Proirity),
		User_uuid:   user,
	}
	return s.repo.CreateReport(report)
}
