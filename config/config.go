package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

var AppConfig = &Config{}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unable to decode config, %v", err)
	}

	InitDB()

}
