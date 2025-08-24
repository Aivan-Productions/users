package configuration

import (
	"fmt"
	"os"
	"slices"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	PORT       int    `env:"PORT" envDefault:"8080"`
	LOG_FORMAT string `env:"LOG_FORMAT" envDefault:"text"`
}

func defaultConfig() *Config {
	return &Config{
		PORT:       8080,
		LOG_FORMAT: "text",
	}
}

func validate(cfg *Config) {
	if cfg.PORT <= 0 || cfg.PORT > 65535 {
		cfg.PORT = 8080
	}
	if !slices.Contains([]string{"text", "json"}, cfg.LOG_FORMAT) {
		cfg.LOG_FORMAT = "text"
	}
}

func loadVarToEnv(configFileName string) {
	if _, err := os.Stat(configFileName); err != nil {
		panic(fmt.Sprintf("configuration: file %s not found: %v", configFileName, err))
	}

	if err := godotenv.Load(configFileName); err != nil {
		panic(fmt.Sprintf("configuration: failed to load %s: %v", configFileName, err))
	}
}

func LoadConfig(configFileName string) *Config {
	loadVarToEnv(".env")
	loadVarToEnv(configFileName)

	cfg := defaultConfig()

	if err := env.Parse(cfg); err != nil {
		return cfg
	}

	validate(cfg)

	return cfg
}
