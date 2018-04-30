package weather

// Struct for handling the api response.
type weatherInfo struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

// goWeather queries openweathermap.org api for current weather conditions.
// Use static json file for testing purposes.
func goWeather(city string) (string, error) {
	// TODO: Where all the work happens.
	// 1. Query the json data
	// 2. Return goWeather phrase
}

// makeURI returns a URI that can be used to make the actual API call.
func makeURI(city string) string {
	// apiKey = "bd5e378503939ddaee76f12ad7a97608"
	// Example API call: "http://api.openweathermap.org/data/2.5/weather?q=London&units=imperial&appid="

	// TODO return URI string
}

func getJsonWeatherData(file string) []weatherInfo {
	// file location: /golearn/sdd/webapp/weatherTest.json"

	//TODO: Read the json file and save data into a weatherinfo slice.
}
