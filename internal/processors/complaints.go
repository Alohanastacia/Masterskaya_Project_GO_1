package processors

import "complaint_service/internal/model"

type ComplaintsRepository interface {
	//имплиментируются методы из repository
	GetComplaint(uuid string) (model.GetComplaint, error)
}

type ComplaintsProcessor struct {
	complaintsRepository ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{complaintsRepository}
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
func (p ComplaintsProcessor) GetComplaint(uuid string) (model.GetComplaint, error) {
	return p.complaintsRepository.GetComplaint(uuid)
}
