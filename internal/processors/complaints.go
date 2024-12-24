package processors

import (
	"complaint_service/internal/entity"
	"complaint_service/internal/repository"
	"fmt"
)

type ComplaintsRepository interface {
	FindUsers(UserUUID string) ([]*entity.Users, error)
	//имплиментируются методы из repository
	CreateComplaints(c entity.CreateComplaint) (int64, error)
}

type ComplaintsProcessor struct {
	Authorization
	ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository *repository.ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{
		Authorization: NewAuthService(complaintsRepository.Authorization),
	}
}

func (p *ComplaintsProcessor) CreateComplaints(c entity.CreateComplaint) (int64, error) {

	if c.Category == "" || c.Description == "" || c.Priority == "" {
		return 0, fmt.Errorf("fields are not filled in")
	}

	return p.ComplaintsRepository.CreateComplaints(c)
}

func (p *ComplaintsProcessor) FindUsers(UserUUID string) (entity.Users, error) {
	return p.FindUsers(UserUUID)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
