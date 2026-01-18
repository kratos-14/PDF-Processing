package utils

import (
	"encoding/base64"
	"os"
)

func LookupEnv(key string) (string, bool) {
	// Placeholder for actual environment variable lookup logic
	envVal := os.Getenv(key)
	if envVal != "" {
		return envVal, true
	}
	return "", false
}

func Base64Decode(txt string) (string, error) {
	valBytes, err := base64.StdEncoding.DecodeString(txt)
	if err != nil {
		return "", err
	}
	return string(valBytes), nil
}

func GetEnv(key, defaultValue string) string {
	value, exists := LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
