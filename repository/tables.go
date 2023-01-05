package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
	)
	if err != nil {
		zap.S().Fatalf("Cannot migrate DB: %v", err)
	}
}

type User struct {
	gorm.Model
}
