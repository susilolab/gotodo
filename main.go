package main

import (
	"github.com/gin-gonic/gin"
	"github.com/susilolab/gotodo/services"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	services.TodoServices(r)

	r.Run()
}
