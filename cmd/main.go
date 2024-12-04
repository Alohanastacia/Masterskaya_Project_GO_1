package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

const port = ":8080"

func main() {
	/*
		Инициализируем БД. И коннект прокидываем в CreateComplaintsRepository
		complaintsRepository := repository.CreateComplaintsRepository(db)

		Инициализируем ComplaintsProcessor где у нас будет бизнес логика
		complaintsProcessor := processors.CreateComplaintsProcessor(complaintsRepository)

		Инициализируем ComplaintsHandler, где у нас будут описаны хендлеры
		complaintsHandler := handlers.CreateComplaintsHandler(complaintsProcessor)
	*/

	app := fiber.New()

	/*
		Подключаем роуты. Прокидываем инициализированные хендлеры complaintsHandler
		routes.Complaints(app, complaintsHandler)
	*/
	log.Println("The server is running")
	if err := app.Listen(port); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}
}
