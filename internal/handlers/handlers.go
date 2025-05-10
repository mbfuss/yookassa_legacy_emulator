package handlers

import (
	"net/http"
	"yookassa_legacy_emulator/internal/business/parser"
)

type Handler struct {
	pr  parser.Parser
	url string
}

func NewHandler(pr parser.Parser, url string) *Handler {
	return &Handler{
		pr:  pr,
		url: url,
	}
}

// RegisterHandlers - регистрации обработчиков HTTP-запросов.
func (h *Handler) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc(registerPreAuth, h.RegisterPreAuth)
	mux.HandleFunc(payments, h.PaymentPage)
	mux.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("templates"))))
}
