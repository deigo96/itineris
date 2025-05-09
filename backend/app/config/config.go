package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string
	ServiceHost string
	ServicePort string
	SecretKey   string
	Environment string
	DBConfig
}

type DBConfig struct {
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
}

func GetConfig() *Config {
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	return &Config{
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceHost: os.Getenv("SERVICE_HOST"),
		ServicePort: os.Getenv("SERVICE_PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),
		Environment: os.Getenv("ENVIRONMENT"),

		DBConfig: DBConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbName:     os.Getenv("DB_NAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}
}
