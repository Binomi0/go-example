package routers

import (
	"mygoapp/handlers"
	"mygoapp/handlers/message"
	"mygoapp/libs/authentication"
	"mygoapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFile("favicon.ico", "frontend/public/favicon.ico")
	r.Static("/assets", "frontend/public")
	r.SetTrustedProxies(nil)
	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")

	app := r.Group("/")
	{
		app.GET("/", authentication.AuthMiddleware(), handlers.GetHome)
		app.GET("/home", handlers.GetLanding)
		app.GET("/login", handlers.GetHome)
		app.POST("/login", handlers.Login)
	}

	v1 := r.Group("/api/v1")
	v1.Use(middlewares.MiddlewareV1())
	v1.Use(authentication.AuthMiddleware())
	{
		v1.GET("/", message.GetMessages)
	}

	v2 := r.Group("/api/v2")
	v2.Use(middlewares.MiddlewareV2())
	v2.Use(authentication.AuthMiddleware())
	{
		v2.GET("/", message.GetMessages)
	}

	return r
}
