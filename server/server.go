package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LaunchServer() error {
	router := gin.Default()
	router.GET("/ping", pong)
	return router.Run()
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
