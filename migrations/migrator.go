package migrations

import (
	"backend-go/features/users/data"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(data.Role{})
	db.AutoMigrate(data.User{})
}
