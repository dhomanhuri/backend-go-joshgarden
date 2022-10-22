package business

import "backend-go/features/sensors"

type SensorRepo struct {
	sensorData sensors.Data
}

func SensorBusiness(data sensors.Data) sensors.Business {
	return &SensorRepo{
		sensorData: data,
	}
}

func (repo *SensorRepo) BusInsert(data sensors.Core) error {
	err := repo.sensorData.DataInsert(data)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SensorRepo) BusGet() (sensors.Core, error) {
	result, err := repo.sensorData.DataGet()
	return result, err
}
