package utils

import (
	"github.com/google/uuid"
)

// GenerateOrderID - генерация уникального идентификатора заказа.
func GenerateOrderID() string {
	return uuid.New().String()
}
