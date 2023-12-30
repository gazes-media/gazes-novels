package database

import (
	"fmt"
	"log"

	"github.com/gazes-media/gazes-novels/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

// DB returns the application's database instance.
// If the database instance has not been initialized, it creates a new connection using
// the configuration settings retrieved from the `config.GetConfig()` function.
//
// Example usage:
//
//	db := DB()
//	// Use db for database operations
func DB() *gorm.DB {
	if databaseInstance == nil {
		// Retrieve database configuration from the application's configuration
		cfg := config.GetConfig()

		// Create the data source name (DSN) for the PostgreSQL connection
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
		)

		// Open a new database connection using GORM
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
		}

		// Set the database instance
		databaseInstance = db
	}

	return databaseInstance
}
