package models

import (
	"time"

	"gorm.io/gorm"
)

type ClientCredentials struct {
	ID            uint           `gorm:"primaryKey"`
	Email         string         `gorm:"size:255;unique;not null"`
	Password      string         `gorm:"size:255;not null"`
	EmailVerified bool           `gorm:"default:false"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
