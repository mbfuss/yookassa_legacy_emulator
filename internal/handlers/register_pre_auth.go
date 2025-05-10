package handlers

import (
	"encoding/json"
	"net/http"
	"yookassa_legacy_emulator/internal/business/utils"
	"yookassa_legacy_emulator/internal/resource"
)

// RegisterPreAuth - обработчик маршрута /payment/rest/registerPreAuth.do.
// Принимает POST-запросы с данными о покупке и возвращает JSON-ответ с orderId и formUrl для дальнейшего редиректа.
func (h *Handler) RegisterPreAuth(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Только POST метод для данного пути", http.StatusMethodNotAllowed)
		return
	}
	err := h.pr.Parse(req)
	if err != nil {
		http.Error(res, "Ошибка при разборе данных: "+err.Error(), http.StatusBadRequest)
		return
	}

	orderID := utils.GenerateOrderID()
	formUrl := h.url + payments + "?orderId=" + orderID
	reqData := resource.ResponsePayload{
		OrderId: orderID,
		FormUrl: formUrl,
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	err = json.NewEncoder(res).Encode(reqData)
	if err != nil {
		http.Error(res, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		return
	}
}
