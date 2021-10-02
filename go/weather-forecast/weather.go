// Package weather contains variables and functions to provide the forecast for a specified city.
package weather

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation describes the current location for the weather forecast.
var CurrentLocation string

// Forecast displays the weather forecast for a specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
