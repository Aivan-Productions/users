package configuration

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT int `env:"PORT" envDefault:"2000"`
}

func defaultConfig() *Config {
	return &Config{
		PORT: 2000,
	}
}

func validate(cfg *Config) {
	if cfg.PORT <= 0 || cfg.PORT > 65535 {
		cfg.PORT = 2000
	}
}

func Load() *Config {
	if _, err := os.Stat(".env"); err != nil {
		panic("configuration: файл .env не найден")
	}

	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Sprintf("configuration: не удалось загрузить .env: %v", err))
	}

	cfg := defaultConfig()

	if err := env.Parse(cfg); err != nil {
		return cfg
	}

	validate(cfg)

	return cfg
}
