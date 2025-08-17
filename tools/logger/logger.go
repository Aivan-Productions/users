package logger

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

// Init инициализирует логгер на основе .env файлов
func Init() *slog.Logger {
	// Определяем .env файл в зависимости от окружения
	envFile := ".env.local"
	if os.Getenv("APP_ENV") == "production" {
		envFile = ".env.production"
	}

	// Загружаем переменные окружения
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: couldn't load %s: %v", envFile, err)
	}

	// Получаем формат логов (по умолчанию text)
	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "" {
		logFormat = "text"
	}

	// Создаём соответствующий обработчик
	var handler slog.Handler
	switch logFormat {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	default: // text
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler)
}
