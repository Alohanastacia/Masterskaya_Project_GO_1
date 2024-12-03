package handlers

type ComplaintsProcessor interface {
	//имплиментируются методы из processors
}

type ComplaintsHandler struct {
	complaintsProcessor ComplaintsProcessor
}

func CreateComplaintsHandler(complaintsProcessor ComplaintsProcessor) *ComplaintsHandler {
	return &ComplaintsHandler{complaintsProcessor}
}

// Ниже будут методы-хендлеры. Вызывают через интерфейс ComplaintsProcessor нужные методы бизнес логики
