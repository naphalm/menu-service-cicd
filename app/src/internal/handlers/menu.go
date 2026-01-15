package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naphalm/menu-service-cicd/internal/db"
)

func GetMenu(c *gin.Context) {
	rows, err := db.DB.Query(`
		SELECT id, name, description, price FROM menu_items
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []map[string]interface{}

	for rows.Next() {
		var id int
		var name, desc string
		var price float64

		rows.Scan(&id, &name, &desc, &price)
		items = append(items, gin.H{
			"id":          id,
			"name":        name,
			"description": desc,
			"price":       price,
		})
	}

	c.JSON(http.StatusOK, items)
}
