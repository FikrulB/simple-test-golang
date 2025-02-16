package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	RedisHost string
	RedisPass string
}

func NewEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		AppPort:   getEnv("APP_PORT", "5000"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASS", ""),
		DBName:    getEnv("DB_NAME", "mydb"),
		RedisHost: getEnv("REDIS_HOST", "localhost:6379"),
		RedisPass: getEnv("REDIS_PASS", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
