package server

import (
	"streaming_app/server/router"
)

func Run() {
	router := router.Create()
	router.Run(":3000")
}
