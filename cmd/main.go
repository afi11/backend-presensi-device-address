package main

import (
	"time"
	"backend_presensi_device_address/pkg/auth"
	"backend_presensi_device_address/pkg/common/db"
	"backend_presensi_device_address/pkg/deviceaddress"
	"backend_presensi_device_address/pkg/divisi"
	"backend_presensi_device_address/pkg/jadwal"
	"backend_presensi_device_address/pkg/presensi"
	"backend_presensi_device_address/pkg/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()

	dbHandler := db.Init(dbUrl)

	users.UserRoutes(router, dbHandler)
	auth.AuthRoutes(router, dbHandler)
	divisi.DivisiRoutes(router, dbHandler)
	jadwal.JadwalRoutes(router, dbHandler)
	deviceaddress.DeviceAddressRoutes(router, dbHandler)
	presensi.PresensiRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

		// Cors Handler
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},
			AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
			  return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		  }))

	router.Run(port)

}
