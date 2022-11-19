package weather

type Request struct {
	City string `json:"city"`
}

type Get struct {
	Count  int    `json:"count"`
	Cities []City `json:"cities"`
}

type AvgTemp struct {
	City      string  `json:"city"`
	CntDay    int     `json:"cnt_day"`
	StartDate int64   `json:"start_date"`
	EndDate   int64   `json:"end_date"`
	Temp      float64 `json:"temp"`
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
