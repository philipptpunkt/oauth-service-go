package models

import (
	"time"
)

type ClientRefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	ClientID  uint      `gorm:"not null;index"`
	Token     string    `gorm:"type:text;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	ClientCredentials ClientCredentials `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
}
