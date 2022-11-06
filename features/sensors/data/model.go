package data

import (
	"gorm.io/gorm"

	"backend-go/features/sensors"
)

type Sensor struct {
	gorm.Model
	Soil       string `json:"soil"`
	WaterLevel string `json:"water_level"`
	Water      string `json:"water"`
}

func (sensor Sensor) ToCore() sensors.Core {
	sensorCore := sensors.Core{
		ID:         int(sensor.ID),
		Soil:       sensor.Soil,
		WaterLevel: sensor.WaterLevel,
		Water:      sensor.Water,
		Time:       sensor.CreatedAt.String(),
	}
	return sensorCore
}
func toCoreList(art []Sensor) []sensors.Core {
	var coreList []sensors.Core
	for _, val := range art {
		coreList = append(coreList, val.ToCore())
	}
	return coreList
}
