// Package utils provides common utility functions used across SynthaMind.
package utils

import "os"

// GetEnv retrieves an environment variable with a fallback default value.
// If the variable is not set, it returns the provided fallback value.
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
