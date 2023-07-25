package models

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast Forecast `json:"forecast"`
}

// 1
type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

// 2
type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

// 2.1
type Condition struct {
	Text string `json:"text"`
}

// 3
type Forecast struct {
	Forecastday []Forecastday `json:"forecastday"`
}

// 3.1
type Forecastday struct {
	Hour []Hour `json:"hour"`
}

// 3.1.1
type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}
