package controllers

import (
	"net/http"
	"streaming_app/server/services/twitch"

	"github.com/gin-gonic/gin"
)

func GetUserByName(c *gin.Context) {
	name := c.Query("name")
	user, err := twitch.GetUserByLogin(name)

	if err.Error == nil {
		c.JSON(http.StatusOK, gin.H{"result": user})
	} else {
		c.JSON(err.Status, gin.H{
			"status":  err.Status,
			"message": err.Message,
			"error":   err.Error.Error(),
		})
	}
}

func GetLiveStreamByUser(c *gin.Context) {
	id := c.Param("id")
	stream, err := twitch.GetLiveStreamByUser(id)

	if err.Error == nil {
		c.JSON(http.StatusOK, gin.H{"result": stream})
	} else {
		c.JSON(err.Status, gin.H{
			"status":  err.Status,
			"message": err.Message,
			"error":   err.Error.Error(),
		})
	}
}
