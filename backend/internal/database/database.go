// Package database manages the database connection and provides utility functions for database operations.
package database

import (
	"fmt"
	"log"

	"github.com/gazes-media/gazes-novels/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

// DB returns the database instance and initializes a connection if it's not already established.
// It connects to the configured PostgreSQL database using the values from the application configuration.
func DB() *gorm.DB {
	if databaseInstance == nil {
		cfg := config.GetConfig()

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		databaseInstance = db
	}

	return databaseInstance
}
