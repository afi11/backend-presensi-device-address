package presensi

import (
	"backend_presensi_device_address/pkg/common/models"
	"backend_presensi_device_address/pkg/common/utils/customtime"
	"backend_presensi_device_address/pkg/common/utils/networkaddress"
	"backend_presensi_device_address/pkg/common/utils/token"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type PresensiInput struct {
	TipePresensi    string `json:"tipe_presensi"`
	JadwalID        int64  `json:"jadwal_id"`
	IPAddressClinet string `json:"ip_address"`
	DeviceAddressID int64  `json:"device_address_id"`
	DivisiID        int64  `json:"divisi_id"`
}

type TimePresensi struct {
	IsAblePresensi   bool   `json:"is_able_presensi"`
	TanggalPresensi  string `json:"tanggal_presensi"`
	TipePresensi     string `json:"tipe_presensi"`
	JadwalID         int64  `json:"jadwal_id"`
	JamAwalPresensi  string `json:"jam_awal_presensi"`
	JamAkhirPresensi string `json:"jam_akhir_presensi"`
}

func (h handler) GetTimePresensi(ctx *gin.Context) {
	userId, _ := token.ExtractTokenID(ctx)

	dateNow := customtime.GetDateNow()
	timeNow := customtime.TimeNowHuman()

	var jadwal models.Jadwal
	var presensi models.Presensi
	var tipePresensi string
	var jamAwalPresensi string
	var jamAkhirPresensi string
	var calculateTime time.Duration
	var isAblePresensi bool

	if err := h.DB.Preload("User.Pegawai.Divisi").Preload(clause.Associations).Where("user_id = ?", userId).
		Where("tanggal = ?", dateNow).First(&jadwal).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"data": "Nothing presence for today"})
		return
	}

	// Cek Presensi Hari Ini
	if err := h.DB.Where("user_id = ?", userId).Where("tanggal_presensi = ?", dateNow).First(&presensi).Error; err != nil {
		tipePresensi = "jam_masuk"
		jamAwalPresensi = jadwal.JamMulaiMasuk
		jamAkhirPresensi = jadwal.JamAkhirMasuk
		calculateTime = customtime.CalculateTime(jadwal.JamMulaiMasuk, timeNow)
		if calculateTime.Milliseconds() < 0 {
			isAblePresensi = true
		} else {
			isAblePresensi = false
		}
	} else {
		tipePresensi = "jam_pulang"
		jamAwalPresensi = jadwal.JamMulaiPulang
		jamAkhirPresensi = jadwal.JamAkhirPulang
		calculateTime = customtime.CalculateTime(jadwal.JamMulaiPulang, timeNow)
		if calculateTime.Milliseconds() < 0 {
			isAblePresensi = true
		} else {
			isAblePresensi = false
		}
	}

	// Cek Presensi Completed
	if presensi.JamMasuk != "" && presensi.JamPulang != "" {
		isAblePresensi = false
		jamAwalPresensi = ""
		jamAkhirPresensi = ""
		tipePresensi = "completed"
	}

	jadwalPresensi := TimePresensi{TipePresensi: tipePresensi, JadwalID: int64(jadwal.ID),
		JamAwalPresensi: jamAwalPresensi, JamAkhirPresensi: jamAkhirPresensi, TanggalPresensi: dateNow,
		IsAblePresensi: isAblePresensi}

	ctx.JSON(http.StatusOK, gin.H{"data": jadwalPresensi})
}

func (h handler) SavePresensi(ctx *gin.Context) {

	// Get IP Address
	fmt.Println(networkaddress.GetClientIPByHeaders(ctx.Request))

	body := PresensiInput{}

	// Get User ID
	userId, _ := token.ExtractTokenID(ctx)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var deviceAddress models.DeviceAddress
	var jadwal models.Jadwal
	var presensi models.Presensi

	// Cek IP Address Valid Or NOT
	if err := h.DB.Where("user_id = ?", userId).Where("ip_address = ?", body.IPAddressClinet).First(&deviceAddress).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	// Cek Jadwal
	if err := h.DB.Where("id = ?", body.JadwalID).Where("user_id = ?", userId).First(&jadwal).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	presensi.DeviceAddressID = int64(deviceAddress.ID)
	presensi.JadwalID = body.JadwalID
	presensi.UserID = int64(userId)
	presensi.DivisiID = body.DivisiID

	if body.TipePresensi == "jam_masuk" {
		presensi.TanggalPresensi = customtime.GetDateNow()
		jamMasuk := customtime.TimeNowHuman()
		calculateTime := customtime.CalculateTime(jadwal.JamAkhirMasuk, jamMasuk)
		if calculateTime.Milliseconds() < 0 {
			telatMasuk := customtime.FormatCorrectTime(calculateTime)
			presensi.TelatMasuk = telatMasuk
			presensi.JamMasuk = jamMasuk
			presensi.IsTelat = 1
		} else {
			presensi.TelatMasuk = "00:00:00"
			presensi.JamMasuk = jamMasuk
			presensi.IsTelat = 0
		}
	} else if body.TipePresensi == "jam_pulang" {
		jamPulang := customtime.TimeNowHuman()
		calculateTime := customtime.CalculateTime(jadwal.JamAkhirPulang, jamPulang)
		if calculateTime.Milliseconds() < 0 {
			telatPulang := fmt.Sprint(int(calculateTime.Hours()), ":", int(calculateTime.Minutes()), ":", int(calculateTime.Seconds()))
			presensi.TelatPulang = telatPulang
			presensi.JamPulang = jamPulang
			presensi.IsTelat = 1
		} else {
			presensi.TelatPulang = "00:00:00"
			presensi.JamPulang = jamPulang
		}
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Type presence is not found!!"})
		return
	}

	tx := h.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Insert Presensi
	if err := tx.Create(&presensi).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// Insert Log Activity
	logActivity := models.LogActivities{TipeAktivitas: "presensis", TableDestinationId: presensi.DeviceAddressID,
		UserId: presensi.UserID}
	if err := tx.Create(&logActivity).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	tx.Commit()

	ctx.JSON(http.StatusCreated, gin.H{"message": "presence is successfull"})
}
