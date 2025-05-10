package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// PaymentPage - обработчик для страницы оплаты.
func (h *Handler) PaymentPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Только GET метод для данного пути", http.StatusMethodNotAllowed)
		return
	}

	orderID := req.URL.Query().Get("orderId")
	if orderID == "" {
		http.Error(res, "orderId не указан", http.StatusBadRequest)
		return
	}

	tmplPath := filepath.Join("templates", "payment.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(res, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	tmpl.Execute(res, struct {
		OrderID   string
		ReturnUrl string
	}{
		OrderID:   orderID,
		ReturnUrl: h.pr.RequestPayload.ReturnUrl + "&orderId=" + orderID,
	})
}
