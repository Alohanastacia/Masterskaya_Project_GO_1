package processors

type ComplaintsRepository interface {
	ComplaintsListAdmin(userId string) (err error)
	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	complaintsRepository ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{complaintsRepository}
}

func (p *ComplaintsProcessor) ComplaintsListAdmin(userId string) (err error) {
	return p.complaintsRepository.ComplaintsListAdmin(userId)
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
