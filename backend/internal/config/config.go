package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTimeZone string

	JwtSecret          string
	JwtExpirationHours int

	Port string

	RedisAddress string
	RedisPort    string
}

var Config AppConfig

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found or could not be loaded: %v. Assuming environment variables are set externally.\n", err)
	}

	Config.DBHost = getEnv("DB_HOST", "localhost")
	Config.DBUser = getEnv("DB_USER", "postgres")
	Config.DBPassword = getEnv("DB_PASSWORD", "")
	Config.DBName = getEnv("DB_NAME", "postgres")
	Config.DBPort = getEnv("DB_PORT", "5432")
	Config.DBSSLMode = getEnv("DB_SSLMODE", "disable")

	tz := getEnv("DB_TIMEZONE", "Africa/Nairobi")
	_, err = time.LoadLocation(tz)
	if err != nil {
		log.Fatalf("Error: Invalid DB_TIMEZONE '%s'. Please provide a valid IANA Time Zone string (e.g., 'America/New_York', 'Africa/Nairobi'). Error: %v", tz, err)
	}
	Config.DBTimeZone = tz

	Config.JwtSecret = getEnv("JWT_SECRET", "")
	if Config.JwtSecret == "" {
		log.Fatal("Error: JWT_SECRET environment variable is not set!")
	}
	jwtExpHours, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24"))
	if err != nil {
		log.Printf("Warning: Invalid JWT_EXPIRATION_HOURS '%s'. Using default 24 hours. Error: %v", os.Getenv("JWT_EXPIRATION_HOURS"), err)
		jwtExpHours = 24
	}
	Config.JwtExpirationHours = jwtExpHours

	Config.Port = getEnv("PORT", "8080")

	Config.RedisAddress = getEnv("REDIS_ADDRESS", "localhost")
	Config.RedisPort = getEnv("REDIS_PORT", "6432")

	fmt.Println("Configuration loaded successfully.")
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
