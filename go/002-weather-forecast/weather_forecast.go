// Package weather provides tools to forecast current weather
// conditions of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current weather location.
var CurrentLocation string

// Forecast takes a city and condition string arguments and returns a string
// describing the current weather condition for the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
