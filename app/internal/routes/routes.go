package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naphalm/menu-service-cicd/app/internal/handlers"
)

func RegisterMenuRoutes(r *gin.Engine) {
	r.GET("/menu", handlers.GetMenu)
}
