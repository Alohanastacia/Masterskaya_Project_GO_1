package processors

type ComplaintsRepository interface {
	//имплиментируются методы из repository
}

type ComplaintsProcessor struct {
	complaintsRepository ComplaintsRepository
}

func CreateComplaintsProcessor(complaintsRepository ComplaintsRepository) *ComplaintsProcessor {
	return &ComplaintsProcessor{complaintsRepository}
}

// Ниже будут методы ComplaintsProcessor, которые реализуют бизнес логику вызываются из хендлеров
// Вызывают методы из repository через интерфейс ComplaintsRepository
