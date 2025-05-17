package database

import (
	"fmt"
	"ginwebapp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(cfg *config.AppConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(func() logger.LogLevel {
			if cfg.Environment == "development" {
				return logger.Info
			}
			return logger.Error
		}()),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	return nil
}