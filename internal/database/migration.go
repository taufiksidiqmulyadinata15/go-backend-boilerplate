package database

import (
	"log"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migration completed")
}
