package config

import (
	"log"

	"github.com/spf13/viper"
)

type TwitchConfig struct {
	Url string `json:"url"` // twitch api base url
	Id  string `json:"id"`  // twitch app Client-ID
}

type ApiConfig struct {
	Version string `json:"version"`
}

type Config struct {
	Twitch TwitchConfig `json:"twitch"`
	Api    ApiConfig    `json:"api"`
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
