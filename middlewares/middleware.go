package middlewares

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Lógica del middleware
        c.Next()
    }
}
