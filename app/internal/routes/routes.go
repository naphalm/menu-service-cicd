package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naphalm/menu-service-cicd/app/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)

	menu := r.Group("/menu")
	{
		menu.GET("", handlers.GetAllMenuItems)
		menu.GET("/:id", handlers.GetMenuItem)
		menu.POST("", handlers.CreateMenuItem)
		menu.PUT("/:id", handlers.UpdateMenuItem)
		menu.DELETE("/:id", handlers.DeleteMenuItem)
	}

	return r
}
