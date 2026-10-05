package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBURL      string
}

func LoadAppConfig() *AppConfig {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &AppConfig{
		Port:       port,
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBURL:      os.Getenv("DATABASE_URL"),
	}
}

func (a *AppConfig) GetDSN() string {
	if a.DBURL != "" {
		return a.DBURL
	}
	if a.DBHost == "" {
		return ""
	}
	sslmode := a.DBSSLMode
	if sslmode == "" {
		sslmode = "disable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		a.DBHost, a.DBUser, a.DBPassword, a.DBName, a.DBPort, sslmode)
}
