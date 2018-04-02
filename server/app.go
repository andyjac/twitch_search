package server

import (
	"twitch_search/server/config"
	"twitch_search/server/router"
)

func Run() {
	config := config.Read()

	router := router.Create()
	router.Run(":" + config.Api.Port)
}
