package processors

import (
	"complaint_service/internal/entity"
)

type ComplaintsRepository interface {
	FindUsers(UserUUID string) ([]*entity.Users, error)
	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	complaintsRepository ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{complaintsRepository}
}

func (p *ComplaintsProcessor) FindUsers(UserUUID string) ([]*entity.Users, error) {
	return p.complaintsRepository.FindUsers(UserUUID)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
