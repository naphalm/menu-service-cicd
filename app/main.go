package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naphalm/menu-service-cicd/app/internal/db"
)

func main() {
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load(".env")
	}

	if err := db.Init(); err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	defer db.Close()

	log.Println("menu-service started")

	r := gin.Default()

	// health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// TODO: add your menu endpoints here

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	// run server in a goroutine so we can handle graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// wait for SIGINT/SIGTERM
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("menu-service shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("menu-service stopped")
}
