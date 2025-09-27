package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	Port        string
	DataFile    string
	LogLevel    string
	CacheMaxAge int
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	config := &Config{
		Port:        getEnvOrDefault("PORT", "8080"),
		DataFile:    getEnvOrDefault("DATA_FILE", "data/data.json"),
		LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
		CacheMaxAge: getEnvIntOrDefault("CACHE_MAX_AGE", 3600), // 1 hour default
	}

	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
