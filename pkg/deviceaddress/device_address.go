package deviceaddress

import (
	"backend_presensi_device_address/pkg/common/models"
	"backend_presensi_device_address/pkg/common/utils/pagination"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type DeviceAddressInput struct {
	IpAddress  string `json:"ip_address" binding:"required"`
	MacAddress string `json:"mac_address" binding:"required"`
	UserID     int64  `json:"user_id" binding:"required"`
}

func (h handler) GetAllAddress(ctx *gin.Context) {
	var deviceAddress []models.DeviceAddress
	var dataPagination pagination.Pagination

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	sort := ctx.Query("sort") + " " + ctx.Query("order")

	dataPagination.Limit = limit
	dataPagination.Page = page
	dataPagination.Sort = sort

	h.DB.Scopes(pagination.PaginateWithPreload("User.Pegawai.Divisi", deviceAddress, &dataPagination, h.DB)).Find(&deviceAddress)
	// No Pagination
	//h.DB.Preload("User.Pegawai.Divisi").Preload(clause.Associations).Find(&deviceAddress)

	dataPagination.Rows = deviceAddress

	ctx.JSON(http.StatusOK, gin.H{"data": dataPagination})
}

func (h handler) GetAddress(ctx *gin.Context) {
	var deviceAddress models.DeviceAddress
	if err := h.DB.Preload("User.Pegawai.Divisi").Preload(clause.Associations).Where("id = ?", ctx.Param("id")).First(&deviceAddress).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": deviceAddress})
}

func (h handler) SaveDeviceAddress(ctx *gin.Context) {
	body := DeviceAddressInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var deviceAddress models.DeviceAddress

	deviceAddress.IpAddress = body.IpAddress
	deviceAddress.UserId = body.UserID

	insertDeviceAddress := h.DB.Create(&deviceAddress)
	if insertDeviceAddress.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": insertDeviceAddress.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "data device address successfull to be added"})
}

func (h handler) UpdateDeviceAddress(ctx *gin.Context) {
	body := DeviceAddressInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var deviceAddress models.DeviceAddress

	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&deviceAddress).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	deviceAddress.IpAddress = body.IpAddress
	deviceAddress.UserId = body.UserID

	updateDeviceAddress := h.DB.Updates(&deviceAddress)
	if updateDeviceAddress.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": updateDeviceAddress.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "data device address successfull to be updated"})

}

func (h handler) DeleteDeviceAddress(ctx *gin.Context) {
	var deviceAddress models.DeviceAddress
	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&deviceAddress).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	deletedDeviceAddress := h.DB.Delete(&deviceAddress)
	if deletedDeviceAddress.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": deletedDeviceAddress.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data device address successfull to be deleted"})
}

func (h handler) ImportDeviceAddress(ctx *gin.Context) {

	file, _ := ctx.FormFile("file")

	dst := "temp/" + file.Filename

	ctx.SaveUploadedFile(file, dst)

	dir, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	fileLocation := filepath.Join(dir, "temp", file.Filename)

	csvFile, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}

	// Read File into a Variable
	lines := csv.NewReader(csvFile)
	// Untuk Custom Format Delimiter, Untuk Default comma
	//lines.Comma = '\t'
	lines.Comma = ';'
	//lines.Comma = ','
	isFirstRow := true
	headerMap := make(map[string]int)
	for {
		// Read row
		record, err := lines.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		// Handle first row case
		if isFirstRow {
			isFirstRow = false

			// Add mapping: Column/property name --> record index
			for i, v := range record {
				headerMap[v] = i
			}

			// Skip next code
			continue
		}

		fmt.Println(record[headerMap["UserID"]])

		user_id, _ := strconv.ParseInt(record[headerMap["UserID"]], 10, 0)
		deviceAddress := models.DeviceAddress{IpAddress: record[headerMap["IPAddress"]],
			UserId: user_id,
		}
		if err := h.DB.Create(&deviceAddress).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "device address successfull to import"})
}
