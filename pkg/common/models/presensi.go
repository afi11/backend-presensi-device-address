package models

import (
	"gorm.io/gorm"
)

type Presensi struct {
	gorm.Model
	JamMasuk        string        `gorm:"size:100;not null;unique" json:"jam_masuk"`
	JamPulang       string        `gorm:"default:TIME;not null;unique" json:"jam_pulang"`
	TelatMasuk      string        `gorm:"default:TIME;not null;unique" json:"telat_masuk"`
	TelatPulang     string        `gorm:"default:TIME;not null;unique" json:"telat_pulang"`
	IsTelat         int           `gorm:"not null" json:"is_telat"`
	UserID          int64         `gorm:"not null" json:"user_id"`
	JadwalID        int64         `gorm:"not null" json:"jadwal_id"`
	DeviceAddressID int64         `gorm:"not null" json:"device_address_id"`
	DivisiID        int64         `gorm:"not null" json:"divisi_id"`
	Divisi          Divisi        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Jadwal          Jadwal        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeviceAddress   DeviceAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
