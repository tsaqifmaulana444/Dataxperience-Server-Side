package main

import (
	"dataxperience-server-side/controllers"
	"dataxperience-server-side/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	api := r.Group("/api/")
	{
		api.GET("/categories", controllers.IndexCategory)
		api.POST("/categories", controllers.CreateCategory)
		api.GET("/categories/:id", controllers.ShowCategories)
		api.PUT("/categories/:id", controllers.UpdateCategory)
		api.DELETE("/categories/:id", controllers.DeleteCategory)

		api.GET("/news", controllers.IndexNews)
		api.POST("/news", controllers.CreateNews)
		api.GET("/news/:id", controllers.ShowNews)
		api.PUT("/news/:id", controllers.UpdateNews)
		api.DELETE("/news/:id", controllers.DeleteNews)

		api.POST("/author", controllers.CreateAuthor)
		api.GET("/author/:id", controllers.ShowAuthor)
		api.PUT("/author/:id", controllers.UpdateAuthor)
		api.DELETE("/author/:id", controllers.DeleteAuthor)

		api.POST("/client", controllers.CreateClient)
		api.GET("/client/:id", controllers.ShowClient)
		api.PUT("/client/:id", controllers.UpdateClient)
		api.DELETE("/client/:id", controllers.DeleteClient)
	}

	r.Run(":5000")
}
