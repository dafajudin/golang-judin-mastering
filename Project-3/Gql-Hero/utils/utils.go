package utils

import (
	"log"
	"github.com/spf13/viper"
)

func GetValue(key string) string {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading config file")
	}

	return viper.GetString(key)
}