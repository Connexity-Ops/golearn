package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"os"
)

// Struct for handling the api response.
type weatherInfo struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func (w weatherInfo) String() string {
	temp := math.Floor(w.Main.Temp)
	switch {
	case temp < 60:
		return fmt.Sprintf("It's %.f degrees outside, It's too cold, grab a coffee and learn some go.\n", temp)
	case temp >= 60 && temp <= 90:
		return fmt.Sprintf("It's %.f degrees outside, It's perfect, take a break from learning go and go outside.\n", temp)
	case temp > 90:
		return fmt.Sprintf("It's %.f degrees outside, It's too hot, turn up the AC and learn go.\n", temp)
	}
	return ""
}

// goWeather queries openweathermap.org api for current weather conditions.
// Use static json file for testing purposes.
func goWeather(city string) (string, error) {
	weather, err := pickWeatherInfo(getJsonWeatherData("weatherTest.json"), city)
	if err != nil {
		return "", err
	}
	return weather.String(), nil
}

func pickWeatherInfo(weatherSlice []weatherInfo, city string) (weatherInfo, error) {
	for _, weather := range weatherSlice {
		if weather.Name == city {
			return weather, nil
		}
	}
	var weather weatherInfo
	return weather, errors.New("Could not find city in weatherSlice")
}

// makeURI returns a URI that can be used to make the actual API call.
func makeURI(city string) string {
	apiKey := "bd5e378503939ddaee76f12ad7a97608"
	baseURL := "http://api.openweathermap.org/data/2.5/weather"

	u, _ := url.Parse(baseURL)
	q := url.Values{}

	q.Set("q", city)
	q.Set("units", "imperial")
	q.Set("appid", apiKey)

	u.RawQuery = q.Encode()
	return u.String()
}

func getJsonWeatherData(path string) []weatherInfo {
	// file location: /golearn/sdd/webapp/weatherTest.json"
	var weatherData []weatherInfo
	file, _ := os.Open(path)
	defer file.Close()
	json.NewDecoder(file).Decode(&weatherData)
	return weatherData
}
