package routes

import (
	"os"

	"github.com/Sadja18/payment-flow-backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Auth, invoices, etc routes here...

	// Seed route - disabled in production
	if os.Getenv("APP_RUN_MODE") != "PROD" {
		router.POST("/internal/seed", func(c *gin.Context) {
			controllers.RunSeedData(c, db)
		})
	}
}
