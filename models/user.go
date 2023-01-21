package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	Role      string `gorm:"size:10;not null;" json:"role"`
	PegawaiID uint32 `gorm:"null" json:"pegawai_id"`
}
