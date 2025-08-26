// Package weather provides tools to Forecast weather in Goblinocus.
package weather

// CurrentCondition describes the current condition of a city.
var CurrentCondition string

// CurrentLocation describes the current city we are analysing.
var CurrentLocation string

// Forecast determines the current weather condition of a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
