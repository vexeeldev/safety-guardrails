package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/luki/safety-guardrails-backend/config"
	"github.com/luki/safety-guardrails-backend/internal/controller"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(cors.New(config.LoadCORSConfig()))

	healthController := controller.NewHealthController()

	r.GET("/health", healthController.Check)
}
