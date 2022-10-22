package migrations

import (
	_sensor "backend-go/features/sensors/data"
	_user "backend-go/features/users/data"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(_sensor.Sensor{})
	db.AutoMigrate(_user.Role{})
	db.AutoMigrate(_user.User{})
}
