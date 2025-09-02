package config

import (
	"os"
	"strconv"
)

type ServiceConfig struct {
	HTTPPort int

	Host     string
}

var Config = ServiceConfig{}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func LoadConfig() {
	httpPortStr := getEnv("HTTP_PORT", "80")
	httpPortInt, err := strconv.Atoi(httpPortStr)
	if err != nil {
		httpPortInt = 80 // fallback default
	}

	host := getEnv("HOST", "0.0.0.0")

	Config = ServiceConfig{
		HTTPPort: httpPortInt,
		Host:     host,
	}
}
