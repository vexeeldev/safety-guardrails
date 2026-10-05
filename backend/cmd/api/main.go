package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luki/safety-guardrails-backend/config"
	"github.com/luki/safety-guardrails-backend/routes"
)

func main() {
	appConfig := config.LoadAppConfig()

	if err := config.InitDB(appConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer config.CloseDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:    ":" + appConfig.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Starting API server on port %s...\n", appConfig.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting gracefully")
}
