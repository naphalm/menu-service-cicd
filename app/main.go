package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/naphalm/menu-service-cicd/app/internal/db"
	"github.com/naphalm/menu-service-cicd/app/internal/routes"
)

func main() {
	db.Init()

	r := gin.Default()
	routes.RegisterMenuRoutes(r)

	log.Println("Menu service running on :3001")
	r.Run(":3001")
}
