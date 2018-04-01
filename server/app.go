package server

import (
	"streaming_app/server/config"
	"streaming_app/server/router"
)

func Run() {
	config := config.Read()

	router := router.Create()
	router.Run(":" + config.Api.Port)
}
