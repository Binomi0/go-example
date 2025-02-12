package main

import (
	"log"
	"mygoapp/routers"
)

func main() {
	r := routers.SetupRouter()
	log.Println("Starting main application")
	log.Fatal(r.Run(":8080"))
}
