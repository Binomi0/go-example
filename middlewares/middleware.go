package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MiddlewareV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lógica del middleware
		log.Println("Executing middleware V1")
		log.Println(c.Request.Header.Get("token"))
		c.Next()
	}
}

func MiddlewareV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lógica del middleware

		log.Println("Executing middleware V2")
		log.Println(c.Request.Header.Get("token"))
		c.Next()
	}
}
