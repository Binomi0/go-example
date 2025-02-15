package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func secureHeaders(c *gin.Context) {
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
	c.Header("Feature-Policy", "none")
	c.Header("Server", "Gin-Gonic/3.2.0") // Optional: Hide the server version
	c.Header("Strict-Transport-Security", "max-age=6307200; includeSubdomains; preload")
	c.Header("X-Frame-Options", "DENY") // Optional: Prevent clickjacking attacks
	// c.Header("Content-Security-Policy", "default-src 'self'; img-src 'self' http://localhost; script-src 'self' http://localhost")
	c.Header("Access-Control-Allow-Origin", "*")

}

func MiddlewareV1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lógica del middleware
		secureHeaders(c)

		log.Println("Executing middleware V1")
		log.Println(c.Request.Header.Get("token"))
		c.Next()
	}
}

func MiddlewareV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lógica del middleware
		secureHeaders(c)

		log.Println("Executing middleware V2")
		log.Println(c.Request.Header.Get("token"))
		c.Next()
	}
}
