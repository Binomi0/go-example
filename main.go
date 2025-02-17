package main

import (
	"log"
	"mygoapp/routers"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // Para cargar todas las variables del .env automáticamente
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	value := os.Getenv("ENVIRONMENT")
	if value == "development" {
		log.Println("Running in development mode")
	} else if value == "production" {
		log.Println("Running in production mode")
	} else {
		log.Println("Running in default mode")
	}
	r := routers.SetupRouter()
	log.Println("Starting main application")
	log.Fatal(r.Run(":8080"))
}
