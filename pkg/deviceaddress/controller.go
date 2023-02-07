package deviceaddress

import (
	"backend_presensi_device_address/pkg/common/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func DeviceAddressRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/device-address")
	routes.Use(middlewares.JwtAuthMiddleware())
	routes.GET("/all-device-address", h.GetAllAddress)
	routes.GET("/get-device-address/:id", h.GetAddress)
	routes.POST("/create-device-address", h.SaveDeviceAddress)
	routes.PUT("/update-device-address/:id", h.UpdateDeviceAddress)
	routes.DELETE("/delete-device-address/:id", h.DeleteDeviceAddress)
	routes.POST("/import-device-address", h.ImportDeviceAddress)
}
