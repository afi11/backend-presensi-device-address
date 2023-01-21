package main

import (
	"backend_presensi_device_address/config"
	"backend_presensi_device_address/database"
	"backend_presensi_device_address/server"
)

func main() {
	config.Init()
	database.InitDB()
	server.Init()
}
