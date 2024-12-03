package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

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
	log.Println("Сервер запущен")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
