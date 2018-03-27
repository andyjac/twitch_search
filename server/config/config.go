package config

import (
	"log"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Type     string `json:"type"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
}

func Read() Config {
	var config Config

	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal(err.Error())
	}

	return config
}
