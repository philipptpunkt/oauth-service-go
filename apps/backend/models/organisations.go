package models

import (
	"time"

	"gorm.io/gorm"
)

type Organisation struct {
	ID          uint                 `gorm:"primaryKey"`
	Name        string               `gorm:"size:255;not null"`
	Description string               `gorm:"type:text"`
	LogoURL     string               `gorm:"type:text"`
	OwnerID     uint                 `gorm:"not null"`
	CreatedAt   time.Time            `gorm:"autoCreateTime"`
	UpdatedAt   time.Time            `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt       `gorm:"index"`
	Owner       ClientCredential     `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	Members     []OrganisationMember `gorm:"foreignKey:OrganisationID"`
}
