package handlers

import (
	"mygoapp/libs/authentication"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func GetHome(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Login", "id": "login"})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home", "id": "root"})
		// c.File("./public/index.html")
	}
}

func Login(c *gin.Context) {
	// Lógica para manejar la subida de un formulario
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar credenciales (en un caso real, esto consultaría una base de datos)
	if loginData.Username != "admin" || loginData.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Generar token JWT
	token, err := authentication.GenerateToken(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.SetCookie(
		"jwt",                       // Nombre de la cookie
		token,                       // Valor de la cookie (JWT)
		int(24*time.Hour.Seconds()), // Expiración en segundos (1 día)
		"/",                         // Rango de la cookie (ruta accesible)
		"",                          // Dominio (vacío usa el mismo dominio)
		true,                        // Solo enviar con HTTPS (true para producción)
		true,                        // Accesible solo por HTTP (HttpOnly, no por JavaScript)
	)

	c.JSON(http.StatusOK, gin.H{})
}
