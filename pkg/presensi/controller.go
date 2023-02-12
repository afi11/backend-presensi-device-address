package presensi

import (
	"backend_presensi_device_address/pkg/common/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func PresensiRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/presensi")
	routes.Use(middlewares.JwtAuthMiddleware())
	routes.POST("/save-presensi", h.SavePresensi)
}
