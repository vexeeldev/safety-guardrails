package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/luki/safety-guardrails-backend/models"
)

var DB *gorm.DB

func InitDB(appConfig *AppConfig) error {
	dsn := appConfig.GetDSN()
	if dsn == "" {
		log.Println("Database configuration (DB_HOST/DATABASE_URL) not found, skipping database connection.")
		return nil
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Device{})
	if err != nil {
		return fmt.Errorf("failed to auto migrate models: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Println("Database connected successfully")
	return nil
}

func CloseDB() error {
	if DB == nil {
		return nil
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
