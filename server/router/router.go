package router

import (
	"streaming_app/server/config"
	"streaming_app/server/controllers"

	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	config := config.Read()
	router := gin.Default()

	api := router.Group("/api/" + config.Api.Version)

	api.GET("/users", controllers.GetUserByName)
	api.GET("/streams/:id", controllers.GetLiveStreamByUser)

	return router
}
