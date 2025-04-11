package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppViper = viper.New()

func init() {
	viperPath, err := os.Getwd()
	if err != nil {
		panic("Cannot get project path")
	}
	AppViper.AddConfigPath(viperPath + string(os.PathSeparator) + "config")
	AppViper.SetConfigName("app")
	AppViper.SetConfigType("yaml")

	if err := AppViper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	} else {
		log.Printf("Using config file: %s", AppViper.ConfigFileUsed())
	}
}

func GetString(key string) string {
	return AppViper.GetString(key)
}

func GetInt(key string) int {
	return AppViper.GetInt(key)
}
