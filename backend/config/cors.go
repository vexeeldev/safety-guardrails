package config

import (
	"time"

	"github.com/gin-contrib/cors"
)

func LoadCORSConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.MaxAge = 12 * time.Hour
	return config
}
