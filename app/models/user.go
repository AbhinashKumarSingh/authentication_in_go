package models

import "time"

type Users struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`

	CreatedAt time.Time `json:"created_at"`
}

func (Users) TableName() string {
	return "users"
}
