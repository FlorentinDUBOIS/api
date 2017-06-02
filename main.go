package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/FlorentinDUBOIS/api/controllers"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	user := api.Group("/user")

	user.Use()
	{
		user.GET("/", controllers.FindUsers())
		user.POST("/", controllers.CreateUser())
		user.GET("/:uid", controllers.FindUser())
		user.PUT("/:uid", controllers.UpdateUser())
		user.DELETE("/:uid", controllers.DeleteUser())
	}

	router.Run()
}
