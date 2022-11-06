package dto

import "backend-go/features/sensors"

type SensorRequest struct {
	Soil       string `json:"soil" form:"soil" binding:"required"`
	WaterLevel string `json:"water_level" form:"water_level" binding:"required"`
	Water      string `json:"water" form:"water" binding:"required"`
}

type SensorResponse struct {
	Soil       string `json:"soil" form:"soil" binding:"required"`
	WaterLevel string `json:"water_level" form:"water_level" binding:"required"`
	Water      string `json:"water" form:"water" binding:"required"`
	Time       string `json:"created_at" form:"created_at" binding:"time"`
}

func FromCoreList(data []sensors.Core) []SensorResponse {
	res := []SensorResponse{}
	for key := range data {
		res = append(res, FromCore(data[key]))
	}
	return res
}

func FromCore(sensorcore sensors.Core) SensorResponse {
	SensorResponse := SensorResponse{
		Soil:       sensorcore.Soil,
		WaterLevel: sensorcore.WaterLevel,
		Water:      sensorcore.Water,
		Time:       sensorcore.Time,
	}
	return SensorResponse
}
