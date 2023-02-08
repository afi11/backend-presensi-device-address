package db

import (
	"backend_presensi_device_address/pkg/common/models"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Divisi{}, &models.Pegawai{}, &models.Jadwal{},
		&models.DeviceAddress{},
		&models.TableDestinations{}, &models.LogActivities{},
		&models.Presensi{})

	// Initial Divisi
	if err = db.AutoMigrate(&models.Divisi{}); err == nil && db.Migrator().HasTable(&models.Divisi{}) {
		if err := db.First(&models.Divisi{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.Divisi{NamaDivisi: "Super Admin"})
		}
	}

	// Activity Type
	if err = db.AutoMigrate(&models.TableDestinations{}); err == nil && db.Migrator().HasTable(&models.TableDestinations{}) {
		if err := db.First(&models.TableDestinations{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&models.TableDestinations{Table: "divisis"})
			db.Create(&models.TableDestinations{Table: "device_addresses"})
			db.Create(&models.TableDestinations{Table: "users"})
			db.Create(&models.TableDestinations{Table: "jadwals"})
			db.Create(&models.TableDestinations{Table: "presensis"})
		}
	}

	return db
}
