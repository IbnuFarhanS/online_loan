package model

import "time"

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    User      `json:"id_user"`
	ProductID Product   `json:"id_product"`
	Amount    float64   `json:"amount"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DueDate   time.Time `json:"due_date"`

	Total               float64 `json:"total,omitempty"`
	TotalMonthlyPayment float64 `json:"total_monthly_payment,omitempty"`
	TotalFee            float64 `json:"total_fee,omitempty"`
}
