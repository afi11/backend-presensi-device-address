package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `gorm:"size:255;not null;unique" json:"username"`
	Password  string  `gorm:"size:100;not null;" json:"password"`
	Role      string  `gorm:"size:10;not null;" json:"role"`
	PegawaiID uint    `gorm:"null" json:"pegawai_id"`
	Pegawai   Pegawai `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type DataUser struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	NIK       string `json:"nik"`
	Nama      string `json:"nama"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Alamat    string `json:"alamat"`
	DivisiID  uint32 `json:"divisi_id"`
}
