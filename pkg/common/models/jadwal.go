package models

import "gorm.io/gorm"

type Jadwal struct {
	gorm.Model
	Tanggal        string `gorm:"default:DATE" json:"tanggal"`
	JamMulaiMasuk  string `gorm:"default:TIME" json:"jam_mulai_masuk"`
	JamAkhirMasuk  string `gorm:"default:TIME" json:"jam_akhir_masuk"`
	JamMulaiPulang string `gorm:"default:TIME" json:"jam_mulai_pulang"`
	JamAkhirPulang string `gorm:"default:TIME" json:"jam_akhir_pulang"`
	UserId         int64  `gorm:"null" json:"user_id"`
}
