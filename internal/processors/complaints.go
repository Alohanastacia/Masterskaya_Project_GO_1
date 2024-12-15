package processors

import "complaint_service/internal/repository"

type ComplaintsRepository interface {
	ComplaintsListAdmin(UserUUID string) (repository.Users, error)
	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	complaintsRepository ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{complaintsRepository}
}

func (p *ComplaintsProcessor) ComplaintsListAdmin(UserUUID string) (repository.Users, error) {
	return p.complaintsRepository.ComplaintsListAdmin(UserUUID)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
