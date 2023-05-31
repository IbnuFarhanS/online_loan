package model

import "time"

type AcceptStatus struct {
	ID            int64       `json:"id"`
	TransactionID Transaction `json:"id_transaction"`
	Status        bool        `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
}
