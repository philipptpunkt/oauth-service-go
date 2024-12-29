package models

import (
	"time"

	"gorm.io/gorm"
)

type ClientProfile struct {
	ID                  uint           `gorm:"primaryKey"`
	ClientCredentialsID uint           `gorm:"not null;uniqueIndex"`
	FirstName           string         `gorm:"size:100"`
	LastName            string         `gorm:"size:100"`
	Organisation        string         `gorm:"size:255"`
	JobTitle            string         `gorm:"size:100"`
	ProfilePicture      string         `gorm:"type:text"`
	TimeZone            string         `gorm:"size:50"`
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
