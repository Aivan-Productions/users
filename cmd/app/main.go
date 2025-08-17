package main

import (
	"fmt"

	"github.com/Aivan-Productions/users/tools/configuration"
	"github.com/Aivan-Productions/users/tools/logger"
)

func main() {

	cfg := configuration.Load()

	fmt.Print(cfg.PORT)
	log := logger.Init()

	log.Info("Запуск приложения")
	log.Error("Ошибка подключения к БД", "error", "timeout")
}
