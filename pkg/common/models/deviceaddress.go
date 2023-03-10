package models

import (
	"gorm.io/gorm"
)

type DeviceAddress struct {
	gorm.Model
	IpAddress string `gorm:"size:100;not null;unique" json:"ip_address"`
	UserId    int64  `gorm:"not null" json:"user_id"`
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
