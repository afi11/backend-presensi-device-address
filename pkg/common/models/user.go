package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	Role      string `gorm:"size:10;not null;" json:"role"`
	PegawaiID uint   `gorm:"null" json:"pegawai_id"`
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

//func (u *DataUser) BeforeSave() error {
//turn password into hash
// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// if err != nil {
// 	return err
// }
// u.Password = string(hashedPassword)

//remove spaces in username
// 	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

// 	return nil
// }

// func (u *DataUser) SaveUserAndEmployee(db *gorm.DB) (*DataUser, error) {
// 	db.Transaction(func(tx *gorm.DB) error {

// 		pegawai := Pegawai{NIK: u.NIK, Nama: u.Nama, Email: u.Email, Telephone: u.Telephone, Alamat: u.Alamat, DivisiID: u.DivisiID}
// 		insertPegawai := tx.Create(&pegawai)
// 		if insertPegawai.Error != nil {
// 			return insertPegawai.Error
// 		}

// 		user := User{Username: u.Username, Password: u.Password, Role: "pegawai", PegawaiID: pegawai.ID}
// 		insertUser := tx.Create(&user)
// 		if insertUser.Error != nil {
// 			return insertUser.Error
// 		}

// 		return nil
// 	})

// 	return u, nil
// }
