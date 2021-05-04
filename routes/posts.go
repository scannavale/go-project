package routes

import (
	"github.com/gin-gonic/gin"
	"go-project/controllers"
)

func PostsRoutes(app *gin.Engine) {
	router := app.Group("/api/posts")
	{
		router.GET("/", controllers.GetAllPosts)
		router.POST("/", controllers.CreatePost)
		router.PUT("/:id", controllers.UpdatePost)
		router.DELETE("/:id", controllers.DeletePost)
	}
}