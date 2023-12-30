// Package config provides functionality for retrieving configuration settings.
package config

import (
	"log"

	"github.com/gazes-media/gazes-novels/internal/utils"
)

// Config represents the configuration settings for the application.
type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

var configInstance *Config

// GetConfig retrieves the application configuration settings.
// It uses environment variables (DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
// to initialize the configuration settings. If the configuration has already been
// retrieved, it returns the existing configuration instance.
//
// Example usage:
//
//	config := GetConfig()
//	fmt.Printf("DB Host: %s\n", config.DBHost)
//	fmt.Printf("DB User: %s\n", config.DBUser)
//	fmt.Printf("DB Password: %s\n", config.DBPassword)
//	fmt.Printf("DB Name: %s\n", config.DBName)
//	fmt.Printf("DB Port: %s\n", config.DBPort)
func GetConfig() *Config {
	if configInstance == nil {
		// Retrieve environment variables for database configuration
		v, err := utils.Getenv([]string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"})

		if err != nil {
			log.Fatal(err)
		}

		// Initialize the configuration instance with retrieved values
		configInstance = &Config{
			DBHost:     v["DB_HOST"],
			DBUser:     v["DB_USER"],
			DBPassword: v["DB_PASSWORD"],
			DBName:     v["DB_NAME"],
			DBPort:     v["DB_PORT"],
		}
	}

	return configInstance
}
