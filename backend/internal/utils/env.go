// Package utils provides utility functions for common tasks.
package utils

import (
	"fmt"
	"os"
)

// Getenv retrieves values for the specified environment variable keys.
// It returns a map containing the key-value pairs of the found environment variables.
// An error is returned if any of the specified keys are not found in the environment.
//
// Example usage:
//
//	keys := []string{"KEY1", "KEY2", "KEY3"}
//	values, err := Getenv(keys)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Access values using keys
//	valueForKey1 := values["KEY1"]
//
// If any of the specified environment variables is not found, the error message
// will indicate the missing keys.
//
// Example error:
//
//	missing environment variables: [KEY2 KEY3]
func Getenv(keys []string) (map[string]string, error) {
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
