package models

import "gorm.io/gorm"

type LogActivities struct {
	gorm.Model
	TipeAktivitas      string            `gorm:"not null" json:"tipe_aktivitas"`
	TableDestinationId int64             `gorm:"not null" json:"table_destination_id"`
	UserId             int64             `gorm:"not null" json:"user_id"`
	TableDestination   TableDestinations `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User               User              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
