package controllers

import (
	"net/http"
	"streaming_app/server/models"

	"github.com/gin-gonic/gin"
)

func GetStreamers(c *gin.Context) {
	streamers, err := models.GetStreamers()

	if err.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"result": streamers,
			"count":  len(streamers),
		})
	} else {
		c.JSON(err.Status, gin.H{"error": err})
	}
}

func GetStreamerById(c *gin.Context) {
	id := c.Param("id")
	streamer, err := models.GetStreamerById(id)

	if err.Message == "" {
		c.JSON(http.StatusOK, gin.H{
			"result": streamer,
			"count":  1,
		})
	} else {
		c.JSON(err.Status, gin.H{"error": err})
	}
}
