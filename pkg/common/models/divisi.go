package models

import "gorm.io/gorm"

type Divisi struct {
	gorm.Model
	NamaDivisi string `gorm:"size:20;not null;unique" json:"divisi_nama"`
}
