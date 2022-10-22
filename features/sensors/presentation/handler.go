package presentation

import (
	"backend-go/features/sensors"
	"backend-go/helper"

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
