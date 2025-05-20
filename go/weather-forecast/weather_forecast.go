// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current weather condition for a Goblinocus city. Can be used to forecast weather.
var CurrentCondition string

// CurrentLocation represents the current Goblinocus location for which condition is stored. Can be used to forecast weather.
var CurrentLocation string

// Forecast returns the weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
