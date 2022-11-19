package weather

type Request struct {
	City string `json:"city"`
}

type Get struct {
	Count  int    `json:"count"`
	Cities []City `json:"cities"`
}

type City struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type Temperature struct {
	CityName string  `json:"city_name"`
	Temp     float64 `json:"temp"`
	Dt       int64   `json:"dt"`
}
