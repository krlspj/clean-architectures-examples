package process_transaction

type TransactionDtoInput struct {
	ID                        string  `json:"id"`
	AccountID                 string  `json:"account_id"`
	Amount                    float64 `json:"amount"`
	CreditCardNumber          string  `json:"credit_card_number"`
	CreditCardName            string  `json:"credit_card_name"`
	CreditCardExpirationMonth int     `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  int     `json:"credit_card_expiration_year"`
	CreditCardCVV             int     `json:"credit_card_cvv"`
}

// {"id": "123","account_id": "1","credit_card_number": "40000000000000000","credit_card_name": "Samuel Terra","credit_card_expiration_month": 12,"credit_card_expiration_year": 2024,"credit_card_expiration_cvv": 123,"amount": 1200}
// {"id": "123","account_id": "1","credit_card_number": "4193523830170205","credit_card_name": "Samuel Terra","credit_card_expiration_month": 12,"credit_card_expiration_year": 2024,"credit_card_expiration_cvv": 123,"amount": 1200}
// {"id": "123","account_id": "1","credit_card_number": "4193523830170205","credit_card_name": "Samuel Terra","credit_card_expiration_month": 12,"credit_card_expiration_year": 2024,"credit_card_expiration_cvv": 123,"amount": 900}

type TransactionDtoOutput struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
