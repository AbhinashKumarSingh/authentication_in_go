package models

import "time"

type RefreshTokens struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	Token     string `gorm:"unique"`
	Revoked   bool   `gorm:"default:false"`
	ExpiresAt time.Time
}

func (RefreshTokens) TableName() string {
	return "refresh_tokens"
}
