package dto

type SensorRequest struct {
	Soil       string `json:"soil" form:"soil" binding:"required"`
	WaterLevel string `json:"water_level" form:"water_level" binding:"required"`
	Water      string `json:"water" form:"water" binding:"required"`
}

type SensorResponse struct {
	Soil       string `json:"soil" form:"soil" binding:"required"`
	WaterLevel string `json:"water_level" form:"water_level" binding:"required"`
	Water      string `json:"water" form:"water" binding:"required"`
}
