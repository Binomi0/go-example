package authentication

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

// Estructura para los claims del token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Función para generar un JWT
func GenerateToken(username string) (string, error) {
	secretKey = []byte(os.Getenv("SECRET_KEY"))
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)), // Expira en 1 hora
			Issuer:    "onrubia.es",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// Middleware de autenticación
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tokenString := c.GetHeader("Authorization")
		tokenString, err := c.Cookie("jwt")

		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "login")
			// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// c.Abort()
			return
		}

		if tokenString == "" {
			c.Redirect(http.StatusFound, "/login")

			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			// c.Abort()
			// c.Redirect(401, "login")
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.SetCookie("token", "", -1, "/", "", true, false)
			c.Redirect(http.StatusFound, "/login")

			// c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			// c.Abort()
			// c.Redirect(401, "login")
			return
		}

		// Guardar los claims en el contexto para su uso en los handlers
		c.Set("username", claims.Username)
		log.Println(("Autenticación exitosa"))
		c.Next()
	}
}
