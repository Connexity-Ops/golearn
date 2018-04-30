package weather

import (
	"testing"
)

func Test_weather(t *testing.T) {
	type args struct {
		city string
	}

	tests := []struct {
		name           string
		input          string
		expectedURI    string
		expectedString string
	}{
		{
			name:           "Too cold for go.",
			input:          "Moscow",
			expectedURI:    "http://api.openweathermap.org/data/2.5/weather?q=Moscow&units=imperial&appid=bd5e378503939ddaee76f12ad7a97608",
			expectedString: "It's 45 degrees outside, It's too cold, grab a coffee and learn some go.\n",
		},
		{
			name:           "Nice day to learn some go.",
			input:          "Santa Monica",
			expectedURI:    "http://api.openweathermap.org/data/2.5/weather?q=Santa%20Monica&units=imperial&appid=bd5e378503939ddaee76f12ad7a97608",
			expectedString: "It's 68 degrees outside, It's perfect, take a break from learning go and go outside.\n",
		},
		{
			name:           "Perfect here.",
			input:          "Tahiti",
			expectedURI:    "http://api.openweathermap.org/data/2.5/weather?q=Tahiti&units=imperial&appid=bd5e378503939ddaee76f12ad7a97608",
			expectedString: "It's 82 degrees outside, It's perfect, take a break from learning go and go outside.\n",
		},
		{
			name:           "Too hot here.",
			input:          "Phoenix",
			expectedURI:    "http://api.openweathermap.org/data/2.5/weather?q=Phoenix&units=imperial&appid=bd5e378503939ddaee76f12ad7a97608",
			expectedString: "It's 82 degrees outside, It's too hot, turn up the AC and learn go.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeURI(tt.input); got != tt.expectedURI {
				t.Errorf("makeURI() = %v, want %v", got, tt.expectedURI)
			}

			if got1, _ := goWeather(tt.input); got1 != tt.expectedString {
				t.Errorf("goWeather() = %v, want %v", got1, tt.expectedString)
			}
		})
	}
}
