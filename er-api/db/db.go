package db

import (
	"er-api/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func connectionKey() string {
	POSTGRES_HOST := getEnv("POSTGRES_HOST", "localhost")
	POSTGRES_USER := getEnv("POSTGRES_USER", "postgres")
	POSTGRES_PASSWORD := getEnv("POSTGRES_PASSWORD", "postgres")
	POSTGRES_DB := getEnv("POSTGRES_DB", "postgres")
	POSTGRES_PORT := getEnv("POSTGRES_PORT", "5432")
	POSTGRES_SSLMODE := getEnv("POSTGRES_SSLMODE", "disable")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB, POSTGRES_PORT, POSTGRES_SSLMODE)
}

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionKey()), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Exchanges{})
	fmt.Println("Successfully connected to database!")
	return db
}
