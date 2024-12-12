package main

import (
	"complaint_service/internal/config"
	l "complaint_service/internal/logger"
	"log/slog"

	"github.com/gofiber/fiber"
)

const port = ":8080"

func main() {
	/*
		Далее передаем в наши ручки log *slog.Logger и с ним работаем.
		Для удобства, в каждой ручке можно использовать такую конструкцию, чтоб дальше подтягивалась информация.
		op:="handlers.GetUser"
		log := slogger.Log.log.With(
			slog.String("где вылезла ошибка", op),
		)
	*/
	cfg := config.NewConfig()
	l.SetupLogger(cfg.Env)
	l.Log.Info("Starting project", slog.String("env", cfg.Env))
	l.Log.Debug("debug messages are enabled", slog.String("env", cfg.Env))

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
	l.Log.Info("The server is running", slog.String("port", port))
	if err := app.Listen(port); err != nil {
		l.Log.Error("Server startup error: %v", err)
	}
}
