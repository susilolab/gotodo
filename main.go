package main

import (
    "os"
    "log"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/stnc/pongo2gin"
    "github.com/joho/godotenv"
	"github.com/susilolab/gotodo/models"
	"github.com/susilolab/gotodo/services"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalln(err)
    }
    dbUrl := os.Getenv("DATABASE_URL")
    models.SetupDB(dbUrl)

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.HTMLRender = pongo2gin.TemplatePath("views")
    r.Static("/public", "./public")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", pongo2.Context{
			"title": "Todo App",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.HTML(200, "hello.html", pongo2.Context{
			"title": "Hello pongo",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	services.TodoServices(r)

	r.Run()
}
