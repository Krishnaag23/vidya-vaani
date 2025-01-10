package config

import (
	"os"
)

var (
	Port        string
	DatabaseURL string
)

func LoadConfig() {
	Port = getEnv("PORT", "8080")
	DatabaseURL = getEnv("DATABASE_URL", "postgres://krishnaag23:testingpass@db:5432/vidya?sslmode=disable")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
