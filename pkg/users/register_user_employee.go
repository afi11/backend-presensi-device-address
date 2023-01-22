package users

import (
	"backend_presensi_device_address/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	NIK       string `json:"nik" binding:"required"`
	Nama      string `json:"nama" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Telephone string `json:"telephone" binding:"required"`
	Alamat    string `json:"alamat" binding:"required"`
	DivisiID  uint32 `json:"divisi_id" binding:"required"`
}

func setHashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

func (h handler) RegisterUserEmployee(ctx *gin.Context) {
	body := RegisterInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.DataUser

	user.NIK = body.NIK
	user.Nama = body.Nama
	user.Email = body.Email
	user.Telephone = body.Telephone
	user.Alamat = body.Alamat
	user.DivisiID = body.DivisiID
	user.Username = body.Username
	user.Password = setHashedPassword(body.Password)

	h.DB.Transaction(func(tx *gorm.DB) error {
		pegawai := models.Pegawai{NIK: user.NIK, Nama: user.Nama, Email: user.Email, Telephone: user.Telephone, Alamat: user.Alamat, DivisiID: user.DivisiID}
		insertPegawai := tx.Create(&pegawai)
		if insertPegawai.Error != nil {
			return insertPegawai.Error
		}

		user := models.User{Username: user.Username, Password: user.Password, Role: "pegawai", PegawaiID: pegawai.ID}
		insertUser := tx.Create(&user)
		if insertUser.Error != nil {
			return insertUser.Error
		}

		return nil
	})

	ctx.JSON(http.StatusCreated, gin.H{"message": "registration success"})
}
