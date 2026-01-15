package main

import (
	"log"

	"github.com/naphalm/menu-service-cicd/app/internal/db"
	"github.com/naphalm/menu-service-cicd/app/internal/routes"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	defer db.Close()

	r := routes.SetupRouter()
	log.Println("menu-service running on :3000")
	r.Run(":3000")
}
