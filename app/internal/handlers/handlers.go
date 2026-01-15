package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// health check
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// placeholders
func GetAllMenuItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"items": []string{}})
}

func GetMenuItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"item": c.Param("id")})
}

func CreateMenuItem(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func UpdateMenuItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func DeleteMenuItem(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
