package resource

// ResponsePayload структура ответа отправляемая сервером клиенту.
type ResponsePayload struct {
	OrderId string `json:"orderId"`
	FormUrl string `json:"formUrl"`
}
