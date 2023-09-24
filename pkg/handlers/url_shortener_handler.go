package handlers

import (
	"github.com/gin-gonic/gin"
)

func UrlShortenerHandler(router *gin.Engine) {

	route := router.Group("/v1")

	route.GET("/url", GetUrl)

}

func GetUrl(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})

	return
}
