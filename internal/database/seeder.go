package database

import (
	"log"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64

	db.Model(&model.User{}).Where("email = ?", "admin@example.com").Count(&count)
	if count > 0 {
		log.Println("Admin user already exists")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash admin password: ", err)
	}

	admin := model.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	if err := db.Create(&admin).Error; err != nil {
		log.Fatal("Failed to seed admin user: ", err)
	}

	log.Println("Admin user seeded successfully")
}
