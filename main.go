package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-project/database"
	"go-project/routes"
	"log"
	"net/http"
)

func main() {
	app := gin.Default()

	// Database connection
	database.Connect()

	// Middlewares
	app.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	// Routes
	routes.PostsRoutes(app)
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	log.Fatal(app.Run(":3000"))
}
