// Package config retrieves and manages application configuration settings.
package config

import (
	"log"

	"github.com/gazes-media/gazes-novels/internal/utils"
)

// Config represents the application configuration settings for the database connection.
type Config struct {
	DBHost     string // Hostname of the database server
	DBUser     string // Username for the database connection
	DBPassword string // Password for the database connection
	DBName     string // Name of the database
	DBPort     string // Port number for the database connection
}

var configInstance *Config

// GetConfig retrieves the application configuration settings for the database.
// It fetches environment variables related to database configuration and initializes the Config instance.
func GetConfig() *Config {
	if configInstance == nil {
		v, err := utils.GetEnv([]string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"})

		if err != nil {
			log.Fatal(err)
		}

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
