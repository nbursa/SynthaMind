package utils

import "os"

// GetEnv retrieves environment variables with a fallback default
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
