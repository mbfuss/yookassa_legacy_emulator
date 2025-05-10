package di

import (
	"fmt"
	"net/http"
	"yookassa_legacy_emulator/internal/business/parser"
	"yookassa_legacy_emulator/internal/config"
	"yookassa_legacy_emulator/internal/handlers"
	"yookassa_legacy_emulator/internal/resource"
)

type Container struct {
	Config  *config.Config
	Handler *handlers.Handler
	Mux     *http.ServeMux
}

// NewContainer - создаёт и инициализирует контейнер.
func NewContainer() *Container {
	cfg := (&config.Config{}).Load()
	rq := resource.RequestPayload{}
	pr := parser.NewParser(rq)
	url := "http://" + cfg.ResourceConfig.Host + cfg.ResourceConfig.Port
	handler := handlers.NewHandler(*pr, url)
	mux := http.NewServeMux()

	return &Container{
		Config:  cfg,
		Handler: handler,
		Mux:     mux,
	}
}

// StartServer - запускает HTTP-сервер.
func (c *Container) StartServer() error {
	if c.Mux == nil {
		return fmt.Errorf("mux не инициализирован")
	}
	c.Handler.RegisterHandlers(c.Mux)
	err := http.ListenAndServe(c.Config.ResourceConfig.Port, c.Mux)
	if err != nil {
		return fmt.Errorf("запуск сервера: %v", err)
	}
	return nil
}
