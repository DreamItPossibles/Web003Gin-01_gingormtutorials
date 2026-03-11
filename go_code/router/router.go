package router

import (
	"go_code/controllers"
	"go_code/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("login", controllers.Login)
		auth.POST("register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRate", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchangeRate", controllers.CreateExchangeRate)
	}

	return r
}
