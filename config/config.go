package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigType("yaml")
	config.AddConfigPath("config/")
	config.SetConfigName("app.config")
	config.WatchConfig()

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
