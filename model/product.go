package model

import "time"

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Installment float64   `json:"installment"`
	Interest    float64   `json:"interest"`
	CreatedAt   time.Time `json:"created_at"`
}
