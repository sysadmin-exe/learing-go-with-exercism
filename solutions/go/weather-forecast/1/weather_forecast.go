// Package weather is used to forecast the weather in Goblinocus.
package weather

// CurrentCondition is used to represent the current weather condition.
var CurrentCondition string
// CurrentLocation is used to represent the current location.
var CurrentLocation string

// Forecast takes in city name and condition as a string and returns a string showing the weather condition in that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
