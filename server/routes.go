package server

import (
	"backend_presensi_device_address/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	public := router.Group("/api")
	public.GET("/test", controllers.Test)

	return router
}
