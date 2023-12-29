package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

var config *Config

// GetConfig returns the application configuration.
// It initializes the config if it has not already been set,
func GetConfig() *Config {
	if config == nil {
		config = &Config{}
		config.Port = getEnv("PORT")
	}

	return config
}

// getEnv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// If the environment variable is not set, it logs a fatal error message
// containing the key and terminates the program.
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}
