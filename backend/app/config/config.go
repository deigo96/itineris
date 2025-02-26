package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName string
	ServiceHost string
	ServicePort string
	SecretKey   string
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
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return &Config{
		ServiceName: os.Getenv("SERVICE_NAME"),
		ServiceHost: os.Getenv("SERVICE_HOST"),
		ServicePort: os.Getenv("SERVICE_PORT"),
		SecretKey:   os.Getenv("SECRET_KEY"),

		DBConfig: DBConfig{
			DbHost:     os.Getenv("DB_HOST"),
			DbPort:     os.Getenv("DB_PORT"),
			DbUser:     os.Getenv("DB_USER"),
			DbName:     os.Getenv("DB_NAME"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}
}
