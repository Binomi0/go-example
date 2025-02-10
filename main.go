// package main

// import (
// 	"log"
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	v1 := r.Group("/api/v1")
// 	{
// 		v1.GET("/", handlers.GetHello)
// 	}
// 	log.Fatal(r.Run(":8080"))
// }

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
