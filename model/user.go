package model

import "time"

type User struct {
	ID          int64     `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	NIK         string    `json:"nik"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Limit       float64   `json:"user_limit"`
	RoleID      Role      `json:"id_role"`
	CreatedAt   time.Time `json:"created_at"`
}
