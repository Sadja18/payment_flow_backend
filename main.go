package main

import (
	"log"
	"time"

	"github.com/Sadja18/payment-flow-backend/config"
	"github.com/Sadja18/payment-flow-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, continuing...")
	}
}

func main() {
	db := config.ConnectDB()
	router := gin.Default()

	// CORS config
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"}, // for dev, you can later restrict this
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	routes.RegisterRoutes(router, db)
	router.Run("0.0.0.0:8080")
}
