package jadwal

import (
	"backend_presensi_device_address/pkg/common/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func JadwalRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/jadwal")
	routes.Use(middlewares.JwtAuthMiddleware())
	routes.POST("/create-jadwal", h.SaveJadwal)
	routes.POST("/import-jadwal", h.ImportJadwal)
}
