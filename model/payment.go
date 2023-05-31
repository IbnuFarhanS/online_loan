package model

import "time"

type Payment struct {
	ID              int64         `json:"id"`
	TransactionID   Transaction   `json:"id_transaction"`
	PaymentAmount   float64       `json:"payment_amount"`
	PaymentMethodID PaymentMethod `json:"id_payment_method"`
	PaymentDate     time.Time     `json:"payment_date"`

	NextInstallment float64 `json:"next_installment,omitempty"`
}
