package handlers

import (
    "net/http"
)

func GetHello(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
