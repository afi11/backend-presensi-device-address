package models

import "gorm.io/gorm"

type DeviceAddress struct {
	gorm.Model
	IpAddress  string `gorm:"size:20;not null;unique" json:"ip_address"`
	MacAddress string `gorm:"size:20;not null;unique" json:"mac_address"`
	UserId     int64  `gorm:"not null" json:"user_id"`
	User       User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
