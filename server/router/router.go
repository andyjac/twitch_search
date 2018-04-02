package router

import (
	"twitch_search/server/config"
	"twitch_search/server/controllers"

	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	config := config.Read()
	router := gin.Default()

	api := router.Group("/api/" + config.Api.Version)

	api.GET("/users", controllers.GetUsersByName)
	api.GET("/users/:id/stream", controllers.GetLiveStreamByUserId)

	return router
}
