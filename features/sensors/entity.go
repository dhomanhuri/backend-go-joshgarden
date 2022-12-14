package sensors

type Core struct {
	ID         int
	Soil       string
	WaterLevel string
	Water      string
	Time       string
}

type Data interface {
	DataInsert(data Core) error
	DataGet() (Core, error)
	DataGetList() ([]Core, error)
}

type Business interface {
	BusInsert(data Core) error
	BusGet() (Core, error)
	BusGetList() ([]Core, error)
}
