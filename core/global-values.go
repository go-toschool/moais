package core

// WisebotGlobalValues ...
type WisebotGlobalValues struct {
	MaxAirTemperature float64
	MinAirTemperature float64
	MaxAirHumidity    float64
	MinAirHumidity    float64
	MaxSoilHumidity   float64
	MinSoilHumidity   float64
}

func (gb *WisebotGlobalValues) setMaxAirTemperature(value float64) {
	gb.MaxAirTemperature = value
}

func (gb *WisebotGlobalValues) setMinAirTemperature(value float64) {
	gb.MinAirTemperature = value
}

func (gb *WisebotGlobalValues) setMaxAirHumidity(value float64) {
	gb.MaxAirHumidity = value
}

func (gb *WisebotGlobalValues) setMinAirHumidity(value float64) {
	gb.MinAirHumidity = value
}

func (gb *WisebotGlobalValues) setMaxSoilHumidity(value float64) {
	gb.MaxSoilHumidity = value
}

func (gb *WisebotGlobalValues) setMinSoilHumidity(value float64) {
	gb.MinSoilHumidity = value
}
