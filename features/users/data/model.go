package data

import (
	"backend-go/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
	Role     Role
}

type Role struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	RoleName string `json:"role_name"`
}

// preload
func (user User) ToCore() users.Core {
	userCore := users.Core{
		UserID:   int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     users.RoleCore{ID: user.RoleID},
	}
	return userCore
}
func toCoreList(art []User) []users.Core {
	var coreList []users.Core
	for _, val := range art {
		coreList = append(coreList, val.ToCore())
	}
	return coreList
}
