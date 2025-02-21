package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DbHost,
		config.DbUser,
		config.DbPassword,
		config.DbName,
		config.DbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Database connected")
	return db
}

func DBCloseConnection(db *gorm.DB) {
	if db != nil {
		dbSQL, _ := db.DB()
		dbSQL.Close()
	}
}
