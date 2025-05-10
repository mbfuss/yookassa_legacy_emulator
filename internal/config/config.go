package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"yookassa_legacy_emulator/internal/resource"
)

type Config struct {
	ResourceConfig resource.Config
}

var (
	configPath = "cfg/.env"
)

// Load - возвращает структуру со значениями из конфиг файла.
func (c *Config) Load() *Config {
	err := godotenv.Load(configPath)
	if err != nil {
		fmt.Printf("Ошибка чтения конфигурационного файла: %v\n", err)
	}

	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Printf("Значение переменной окружения PORT: %s\n", port)

	if port == "" {
		fmt.Println("Будет установлен порт по умолчанию: 8080")
		port = ":8080"
	}

	return &Config{
		ResourceConfig: resource.Config{
			Port: port,
			Host: host,
		},
	}
}
