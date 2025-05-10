package resource

// RequestPayload - структура для обработки входящих запросов поступающих серверу от клиента.
type RequestPayload struct {
	OrderNumber string `json:"orderNumber"`
	Amount      string `json:"amount"`
	ReturnUrl   string `json:"returnUrl"`
	FailUrl     string `json:"failUrl"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	JsonParams  struct {
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	CustomerDetails struct {
		Email    *string `json:"email"`
		Phone    string  `json:"phone"`
		FullName string  `json:"fullName"`
	} `json:"customerDetails"`

	CartItems struct {
		Items []struct {
			PositionId int    `json:"positionId"`
			Name       string `json:"name"`
			Quantity   struct {
				Value   interface{} `json:"value"`
				Measure string      `json:"measure"`
			} `json:"quantity"`
			Tax struct {
				TaxType int `json:"taxType"`
			} `json:"tax"`
			ItemPrice      int `json:"itemPrice"`
			ItemAttributes struct {
				Attributes []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"itemAttributes"`
		} `json:"items"`
	} `json:"cartItems"`
	TaxSystem string `json:"taxSystem"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
