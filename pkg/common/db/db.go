package db

import (
	"backend_presensi_device_address/pkg/common/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Divisi{}, &models.Pegawai{}, &models.Jadwal{}, &models.DeviceAddress{})

	return db
}
