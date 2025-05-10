package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"yookassa_legacy_emulator/internal/resource"
)

type Parser struct {
	RequestPayload resource.RequestPayload
}

func NewParser(rq resource.RequestPayload) *Parser {
	return &Parser{
		RequestPayload: rq,
	}
}

// Parse - парсит входящие данные.
func (p *Parser) Parse(req *http.Request) error {
	// Разбиваем на ключ-значение, так как формат входящих данных application/x-www-form-urlencoded.
	if err := req.ParseForm(); err != nil {
		return fmt.Errorf("при разборе формы: %v", err)
	}

	parserData := make(map[string]string)
	for key, value := range req.Form {
		for _, v := range value {
			parserData[key] = v
		}
	}

	p.RequestPayload.OrderNumber = parserData["orderNumber"]
	p.RequestPayload.Amount = parserData["amount"]
	p.RequestPayload.ReturnUrl = parserData["returnUrl"]
	p.RequestPayload.FailUrl = parserData["failUrl"]
	p.RequestPayload.Email = parserData["email"]
	p.RequestPayload.Phone = parserData["phone"]

	err := json.Unmarshal([]byte(parserData["jsonParams"]), &p.RequestPayload.JsonParams)
	if err != nil {
		log.Println("При разборе jsonParams JSON:", err)
		return fmt.Errorf("при разборе jsonParams JSON: %v", err)

	}

	if err := json.Unmarshal([]byte(parserData["orderBundle"]), &p.RequestPayload); err != nil {
		log.Println("При разборе orderBundle JSON:", err)
		return fmt.Errorf("при разборе orderBundle JSON: %v", err)
	}

	p.RequestPayload.TaxSystem = parserData["taxSystem"]
	p.RequestPayload.Username = parserData["username"]
	p.RequestPayload.Password = parserData["password"]

	log.Println("Декодированный JSON:", p.RequestPayload)

	return nil
}
