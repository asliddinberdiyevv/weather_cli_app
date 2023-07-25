package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"weather/models"

	"github.com/fatih/color"
)

func main() {
	query := "Tashkent"

	if len(os.Args) >= 2 {
		query = os.Args[1]
	}

	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=91a9297a38e24e58b0894703232507&q=%s&days=1&aqi=no&alerts=no", query)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	if res.StatusCode == 400 {
		panic("this location not found")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s, %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %s\n", date.Format("15:04"), hour.TempC, hour.Condition.Text)

		changing := hour.TempC
		switch {
		case changing > 40:
			color.Red(message)
		case changing > 30:
			color.Yellow(message)
		case changing > 20:
			color.Blue(message)
		default:
			color.Black(message)
		}
	}
}
