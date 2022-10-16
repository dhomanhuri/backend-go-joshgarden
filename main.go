package main

import (
	"backend-go/config"
	"backend-go/features/users"
	"backend-go/features/users/bussiness"
	"backend-go/features/users/data"
	"backend-go/features/users/presentation"
	"backend-go/migrations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB              = config.InitDB()
	userRepo     users.Data            = data.Repository(db)
	UserBussines users.Bussines        = bussiness.UserBussines(userRepo)
	UserHandler  presentation.UserBuss = presentation.UserHandler(UserBussines)
)

func main() {
	// config.InitDB()
	db := config.InitDB()
	migrations.AutoMigrate(db)
	router := gin.Default()
	router.POST("/register", UserHandler.Register)
	router.POST("/login", UserHandler.Login)
	router.Run(":8080")
}
