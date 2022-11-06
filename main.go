package main

import (
	"backend-go/config"
	_sensors "backend-go/features/sensors"
	_sensorsBus "backend-go/features/sensors/Business"
	_sensorsData "backend-go/features/sensors/data"
	_sensorsPres "backend-go/features/sensors/presentation"
	_users "backend-go/features/users"
	_usersBus "backend-go/features/users/bussiness"
	_usersData "backend-go/features/users/data"
	_usersPres "backend-go/features/users/presentation"
	"backend-go/migrations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB               = config.InitDB()
	userRepo       _users.Data            = _usersData.Repository(db)
	UserBussines   _users.Bussines        = _usersBus.UserBussines(userRepo)
	UserHandler    _usersPres.UserBuss    = _usersPres.UserHandler(UserBussines)
	sensorRepo     _sensors.Data          = _sensorsData.Repository(db)
	sensorBussines _sensors.Business      = _sensorsBus.SensorBusiness(sensorRepo)
	sensorHandler  _sensorsPres.SensorBus = _sensorsPres.SensorHandler(sensorBussines)
)

func main() {
	// config.InitDB()
	db := config.InitDB()
	migrations.AutoMigrate(db)
	router := gin.Default()
	router.POST("/api/register", UserHandler.Register)
	router.POST("/api/login", UserHandler.Login)
	router.GET("/api/profile", UserHandler.Profile)
	router.GET("/api/user", UserHandler.UserAll)
	router.DELETE("/api/user", UserHandler.DellUser)
	router.GET("/api/sensor", sensorHandler.GetLastData)
	router.GET("/api/sensor/add", sensorHandler.InsertData)
	router.GET("/api/sensorlist", sensorHandler.GetListData)
	// port := os.Getenv("PORT")
	router.Run()
}
