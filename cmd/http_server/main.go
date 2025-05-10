package main

import (
	"log"
	"yookassa_legacy_emulator/internal/di"
)

func main() {
	container := di.NewContainer()
	err := container.StartServer()
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v\n", err)
	}
}
