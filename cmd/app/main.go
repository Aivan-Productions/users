package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/Aivan-Productions/users/tools/configuration"
	"github.com/Aivan-Productions/users/tools/logger"
)

func main() {
	envFile := flag.String("env", ".env.local", "path to env file")
	flag.Parse()

	// Загружаем конфигурацию
	cfg := configuration.LoadConfig(*envFile)

	// Инициализируем логгер
	log := logger.Init(cfg)

	// Используем логгер вместо fmt.Println
	log.Info("Application starting",
		"port", cfg.PORT,
		"log_format", cfg.LOG_FORMAT,
		"env_file", *envFile)

	log.Debug("This is debug message")
	log.Warn("This is warning message")
	log.Error("This is error message",
		slog.String("error", "test error"))

	// Пример с контекстом
	requestLog := log.With(
		"request_id", "12345",
		"user_id", "67890",
	)
	requestLog.Info("Request processed successfully")

	fmt.Println("Check console for colored logs!")
}
