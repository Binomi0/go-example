package routers

import (
	"mygoapp/handlers"
	"mygoapp/handlers/message"
	"mygoapp/libs/authentication"
	"mygoapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.MiddlewareV1())
	{
		v1.GET("/", handlers.GetHome)
		v1.POST("/login", handlers.Login)
	}

	v2 := r.Group("/api/v2")
	v2.Use(middlewares.MiddlewareV2())
	v2.Use(authentication.AuthMiddleware())
	{
		v2.GET("/", message.GetMessages)
	}

	return r
}
