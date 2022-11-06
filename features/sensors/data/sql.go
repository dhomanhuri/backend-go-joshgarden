package data

import (
	"backend-go/features/sensors"
	"errors"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) sensors.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) DataInsert(data sensors.Core) error {
	sensorModel := Sensor{
		Soil:       data.Soil,
		WaterLevel: data.WaterLevel,
		Water:      data.Water,
	}
	result := db.DBConn.Create(&sensorModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *MysqlDB) DataGet() (sensorCore sensors.Core, err error) {
	sensorModel := Sensor{}
	result := repo.DBConn.Last(&sensorModel)
	if result.RowsAffected == 0 {
		return sensorCore, errors.New("failed")
	}
	sensorCore = sensors.Core{
		Soil:       sensorModel.Soil,
		WaterLevel: sensorModel.WaterLevel,
		Water:      sensorModel.Water,
	}
	return sensorCore, nil
}

func (repo *MysqlDB) DataGetList() ([]sensors.Core, error) {
	var data []Sensor
	result := repo.DBConn.Order("id desc").Limit(24).Find(&data)
	if result.RowsAffected == 0 {
		return nil, errors.New("get data failed")
	}
	return toCoreList(data), nil
}
