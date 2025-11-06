package utils

import "os"

func LookupEnv(key string) (string, bool) {
	// Placeholder for actual environment variable lookup logic
	envVal := os.Getenv(key)
	if envVal != "" {
		return envVal, true
	}
	return "", false
}

func GetEnv(key, defaultValue string) string {
	value, exists := LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
