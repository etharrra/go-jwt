package main

import (
	"github.com/etharrra/go-jwt/controllers"
	"github.com/etharrra/go-jwt/initializers"
	"github.com/etharrra/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	// Public routes
	publicGroup := r.Group("/")
	publicGroup.Use(middleware.NoAuthorize)
	{
		publicGroup.POST("/signup", controllers.Singup)
		publicGroup.POST("/signin", controllers.Login)
	}

	// private routes
	privateGroup := r.Group("/")
	privateGroup.Use(middleware.Authorize)
	{
		privateGroup.GET("/signout", controllers.Singout)
		privateGroup.GET("/home", controllers.Home)
	}

	r.Run(initializers.ServerAdderss)
}
