package models

import (
	"time"

	"gorm.io/gorm"
)

type OrganisationMember struct {
	ID             uint              `gorm:"primaryKey"`
	OrganisationID uint              `gorm:"not null;index"`
	ClientID       uint              `gorm:"not null;index"`
	Role           Role              `gorm:"size:50;not null"`
	JobTitle       string            `gorm:"size:100"`
	JoinedAt       time.Time         `gorm:"autoCreateTime"`
	CreatedAt      time.Time         `gorm:"autoCreateTime"`
	UpdatedAt      time.Time         `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt    `gorm:"index"`
	Organisation   Organisation      `gorm:"foreignKey:OrganisationID;constraint:OnDelete:CASCADE"`
	Client         ClientCredentials `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
}
