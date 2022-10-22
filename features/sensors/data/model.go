package data

type Sensor struct {
	Id         int    `json:"id"`
	Soil       string `json:"soil"`
	WaterLevel string `json:"water_level"`
	Water      string `json:"water"`
}
