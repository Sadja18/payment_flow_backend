package controllers

import (
	"net/http"

	"github.com/Sadja18/payment-flow-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RunSeedData(c *gin.Context, db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Users already seeded"})
		return
	}

	users := []models.User{
		{Name: "Alice", Username: "payer", Password: "1234", Role: models.RolePayer},
		{Name: "Bob", Username: "payee1", Password: "1234", Role: models.RolePayee},
		{Name: "Charlie", Username: "payee2", Password: "1234", Role: models.RolePayee},
		{Name: "Diana", Username: "payer2", Password: "1234", Role: models.RolePayer},
		{Name: "Eve", Username: "payee3", Password: "1234", Role: models.RolePayee},
	}

	for _, u := range users {
		db.Create(&u)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dummy users seeded"})
}
