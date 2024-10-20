package main

import (
	"eventBooker/db"
	"eventBooker/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
