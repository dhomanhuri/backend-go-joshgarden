package presentation

import (
	"backend-go/features/sensors"
	"backend-go/features/sensors/presentation/dto"
	"backend-go/helper"
	"backend-go/middlewares"

	"github.com/gin-gonic/gin"
)

type SensorBus struct {
	Bus sensors.Business
}

func SensorHandler(ss sensors.Business) SensorBus {
	return SensorBus{
		Bus: ss,
	}
}

func (sb SensorBus) InsertData(c *gin.Context) {
	soil := c.Query("soil")
	level := c.Query("water_level")
	water := c.Query("water")
	core := sensors.Core{
		Soil:       soil,
		WaterLevel: level,
		Water:      water,
	}
	result := sb.Bus.BusInsert(core)
	if result != nil {
		c.JSON(helper.BadRequest(result.Error()))
		c.Abort()
		return
	}
	c.JSON(helper.SiccessCreate())
}

func (sb SensorBus) GetLastData(c *gin.Context) {
	result, err := sb.Bus.BusGet()
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG("tidak ditemukan data"))
		return
	}
	fromCore := sensors.Core{
		Soil:       result.Soil,
		WaterLevel: result.WaterLevel,
		Water:      result.Water,
	}
	c.JSON(helper.SuccessGetData(fromCore))
}

func (sb SensorBus) GetListData(c *gin.Context) {
	_, _, errJWT := middlewares.JWTTokenCheck(c)
	if errJWT != nil {
		c.JSON(helper.FailedBadRequestWithMSG("invalid or exp jwt"))
		return
	}
	result, err := sb.Bus.BusGetList()
	if err != nil {
		c.JSON(helper.FailedBadRequestWithMSG("tidak ditemukan data / "))
		return
	}
	c.JSON(helper.SuccessGetData(dto.FromCoreList(result)))
}
