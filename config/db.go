package config

import (
	"fmt"
	"github/teacoder-team/golang-backend/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
		cfg.PostgresPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to establish database connection: %v", err)
	}

	log.Println("✅ Database connection established successfully")

	err = DB.AutoMigrate(
		&models.User{},
		&models.Course{},
	)
	if err != nil {
		log.Fatalf("❌ Error during migration: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
