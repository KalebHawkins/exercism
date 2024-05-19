// Package weather forecasts the weather for various cities in Goblinocus.
package weather

// CurrentCondition describes the weather's current condition.
var CurrentCondition string

// CurrentLocation describes the city of the current weather forecast.
var CurrentLocation string

// Forecast provides a description of the weather in the provided *city* and *condition*.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
