package server

import "backend_presensi_device_address/config"

func Init() {
	config := config.GetConfig()
	r := Router()
	r.Run(":" + config.GetString("server.port"))
}
