package parser

import "fmt"

func (p *Parser) VerificationParser() error {
	if p.RequestPayload.Username == "" {
		return fmt.Errorf("параметр Username не может быть пустым")
	}

	if p.RequestPayload.Password == "" {
		return fmt.Errorf("параметр Password не может быть пустым")
	}

	if p.RequestPayload.OrderNumber == "" {
		return fmt.Errorf("параметр OrderNumber не может быть пустым")
	}

	if p.RequestPayload.ReturnUrl == "" {
		return fmt.Errorf("параметр ReturnUrl не может быть пустым")
	}

	if p.RequestPayload.FailUrl == "" {
		return fmt.Errorf("параметр FailUrl не может быть пустым")
	}

	if p.RequestPayload.CustomerDetails.Email == "" && p.RequestPayload.CustomerDetails.Phone == "" {
		return fmt.Errorf("в CustomerDetails должны быть указаны email или phone")
	}

	return nil
}
