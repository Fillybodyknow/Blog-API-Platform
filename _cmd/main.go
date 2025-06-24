package main

import (
	"log"
	"os"

	"github.com/Fillybodyknow/blog-api/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := config.ConnectMongo()
	defer client.Disconnect(nil)

	r := gin.Default()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
