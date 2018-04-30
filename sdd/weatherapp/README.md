Weather lookup
======

Weather queries openweathermap.org api for current weather conditions. (but not really)
Weather takes a city name and returns the current temperature with a response about learning go.

// If temp is 60 degrees or less: "It's too cold, grab a coffee and learn some go."
// If temp is more than 93 degrees: "It's too hot, turn up the AC and learn go."
// Everything else: "It's perfect, take a break from learning go and go outside."

The final response should include the temperature in the string.

// Example: "It's 78 degrees outside, It's perfect, take a break from learning go and go outside."

Since the api temp reading is very dynamic I have provided a test json to use instead of the http response.
Unmarshal the returned JSON and print a phrase describing how you should learn go based on the temp. 


## Example
```go
fmt.Println(goWeather("Moscow"))
```
Returns:
```go
"It's 45 degrees outside, It's too cold, grab a coffee and learn some go."
```

## Testing
```
 go test
```

## Notes
The test checks that a correct URI is created even though it does not make the call.
Example: "http://api.openweathermap.org/data/2.5/weather?q=Moscow&units=imperial&appid=bd5e378503939ddaee76f12ad7a97608",
Details are in the stub file.