package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	ServerPort      string
	MongoURI        string
	MongoDatabase   string
	JWTSecret       string
	GoogleClientID  string
	GoogleClientSecret string
}

// Load loads configuration from environment variables
func Load() *Config {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading it. Using environment variables.")
	}

	return &Config{
		ServerPort:      getEnv("SERVER_PORT", "8080"),
		MongoURI:        getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDatabase:   getEnv("MONGO_DATABASE", "link_manager"),
		JWTSecret:       getEnv("JWT_SECRET", "your-super-secret-key"),
		GoogleClientID:  getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
