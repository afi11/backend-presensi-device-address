package divisi

import (
	"backend_presensi_device_address/pkg/common/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func DivisiRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/divisi")
	routes.Use(middlewares.JwtAuthMiddleware())
	routes.GET("/all-divisi", h.GetAllDivisi)
	routes.POST("/create-divisi", h.SaveDivisi)
	routes.GET("/get-divisi/:id", h.GetDivisi)
	routes.PUT("/update-divisi/:id", h.UpdateDivisi)
	routes.DELETE("/delete-divisi/:id", h.DeleteDivisi)
}
