package initializers

import (
	"github.com/etharrra/go-jwt/models"
)

func SyncDatabase() {
	if !DB.Migrator().HasTable(&models.User{}) {
		DB.AutoMigrate(&models.User{})
	}
}
