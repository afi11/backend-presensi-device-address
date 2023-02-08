package models

import "gorm.io/gorm"

type TableDestinations struct {
	gorm.Model
	Table string `gorm:"not null" json:"table"`
}
