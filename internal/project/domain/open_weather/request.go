package open_weather

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type Weather struct {
	Current struct {
		Temp float64 `json:"temp"`
	} `json:"current"`
}
