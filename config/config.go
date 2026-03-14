package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	JWTExpiryHours int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("no .env file found")
	}

	expiryHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_HOURS"))
	if err != nil {
		expiryHours = 24
	}

	return &Config{
		AppPort:        os.Getenv("APP_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTExpiryHours: expiryHours,
	}
}
