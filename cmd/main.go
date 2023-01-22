package main

import (
	"backend_presensi_device_address/pkg/auth"
	"backend_presensi_device_address/pkg/common/db"
	"backend_presensi_device_address/pkg/users"

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

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

	router.Run(port)

}
