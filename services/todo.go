package services

import (
	"github.com/gin-gonic/gin"
	"github.com/susilolab/gotodo/models"
)

func TodoServices(app *gin.Engine) {
	todo := app.Group("/todo")

	// `/todo`
	todo.GET("/", func(c *gin.Context) {
		c.String(200, "Todo")
	})

	// `/todo/1`
	todo.GET("/:id", func(c *gin.Context) {
		c.String(200, "Todo"+c.Param("id"))
	})

	// `/todo`
	todo.POST("/", func(c *gin.Context) {
		var t models.Todo
		if err := c.BindJSON(&t); err != nil {
			return
		}

		c.JSON(200, t)
	})
}
