package divisi

import (
	"backend_presensi_device_address/pkg/common/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DivisiInput struct {
	NamaDivisi string `json:"nama_divisi" binding:"required"`
}

func (h handler) GetAllDivisi(ctx *gin.Context) {
	var divisis []models.Divisi
	h.DB.Find(&divisis)

	ctx.JSON(http.StatusOK, gin.H{"data": divisis})
}

func (h handler) SaveDivisi(ctx *gin.Context) {
	body := DivisiInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var divisi models.Divisi

	divisi.NamaDivisi = body.NamaDivisi

	insertDivisi := h.DB.Create(&divisi)
	if insertDivisi.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": insertDivisi.Error})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "data division successfull to be added"})

}

func (h handler) GetDivisi(ctx *gin.Context) {
	var divisi models.Divisi
	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&divisi).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": divisi})
}

func (h handler) UpdateDivisi(ctx *gin.Context) {
	body := DivisiInput{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var divisi models.Divisi

	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&divisi).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	divisi.NamaDivisi = body.NamaDivisi

	h.DB.Updates(&divisi)
	ctx.JSON(http.StatusOK, gin.H{"message": "data division successfull to be updated"})
}

func (h handler) DeleteDivisi(ctx *gin.Context) {

	var divisi models.Divisi

	if err := h.DB.Where("id = ?", ctx.Param("id")).First(&divisi).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data is not found!!"})
		return
	}

	h.DB.Delete(&divisi)

	ctx.JSON(http.StatusOK, gin.H{"message": "data division successfull to be deleted"})

}
