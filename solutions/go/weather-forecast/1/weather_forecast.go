// Package weather is a package that returns the current weather condition for a given city in Goblinocus.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the city for which the current weather forecast is required.
var CurrentLocation string

// Forecast shows/returns the current weather condition for the selected city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
