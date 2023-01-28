package infrastructure

import (
	"github.com/yagi-eng/go-pj-template2/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(repository.Tables()...); err != nil {
		zap.S().Fatalf("Cannot migrate DB: %v", err)
	}
}
