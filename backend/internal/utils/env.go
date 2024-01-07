// Package utils provides utility functions for handling environment variables.
package utils

import (
	"fmt"
	"os"
)

// Getenv retrieves the values of environment variables specified by the keys.
// It returns a map of key-value pairs of the retrieved environment variable values.
// If any specified environment variable is missing, it returns an error indicating the missing variables.
func GetEnv(keys []string) (map[string]string, error) {
	values := make(map[string]string)
	missing := []string{}

	for i := range keys {
		value := os.Getenv(keys[i])

		if value == "" {
			missing = append(missing, keys[i])
		} else {
			values[keys[i]] = value
		}
	}

	if len(missing) > 0 {
		return values, fmt.Errorf("missing environment variables: %v", missing)
	}

	return values, nil
}
