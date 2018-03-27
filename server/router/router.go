package router

import (
	"streaming_app/server/controllers"

	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	router := gin.Default()

	router.GET("/streamers", controllers.GetStreamers)
	router.GET("/streamers/:id", controllers.GetStreamerById)

	return router
}
