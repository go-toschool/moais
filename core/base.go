package core

var (
	globalValues WisebotGlobalValues
)

func main () {
	globalValues = WisebotGlobalValues{
		MaxAirTemperature = 10
		MinAirTemperature = 10
		MaxAirHumidity    = 10
		MinAirHumidity    = 10
		MaxSoilHumidity   = 10
		MinSoilHumidity   = 10
	}
}