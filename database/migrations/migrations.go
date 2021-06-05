package migrations

import (
	"github.com/repRSilva/api-golang/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Users{})
}
