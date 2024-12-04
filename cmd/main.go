package main

import (
	"complaint_service/internal/config"
	"log/slog"
	"os"

	"github.com/gofiber/fiber"
)

const port = ":8080"
const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	/*
		Далее передаем в наши ручки log *slog.Logger и с ним работаем.
		Для удобства, в каждой ручке можно использовать такую конструкцию, чтоб дальше подтягивалась информация.
		log := log.With(
			slog.String("где вылезла ошибка", op),
		)
	*/
	cfg := config.NewConfig()
	log := setupLogger(cfg.Env)
	log.Info("Starting project", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled", slog.String("env", cfg.Env))
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
	log.Info("The server is running", slog.String("port", port))
	if err := app.Listen(port); err != nil {
		log.Error("Server startup error: %v", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log

}
