package db

import (
	"fmt"
	"log"

	"github.com/imdinnesh/prepaid-card/services/auth/config"
	models "github.com/imdinnesh/prepaid-card/services/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate user table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Auto migrate refresh token table
	err = db.AutoMigrate(&models.RefreshToken{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	fmt.Println("Database connected & migrated successfully")
	return db
}
