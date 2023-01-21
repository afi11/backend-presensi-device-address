package models

import "gorm.io/gorm"

type Pegawai struct {
	gorm.Model
	NIK       string `gorm:"size:20;not null;unique" json:"nik"`
	Nama      string `gorm:"size:100;not null;" json:"nama"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Telephone string `gorm:"size:15;not null;" json:"telephone"`
	Alamat    string `gorm:"size:15;not null;" json:"alamat"`
	DivisiID  uint32 `gorm:"null" json:"divisi_id"`
}
