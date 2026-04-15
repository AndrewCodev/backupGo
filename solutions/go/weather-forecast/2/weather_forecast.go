// Package weather provides functionality to report current weather conditions for a given location.
package weather

var (
    // CurrentCondition represents the current weather condition.
    CurrentCondition string

    // CurrentLocation represents the current location for the weather report.
    CurrentLocation string
)

// Forecast returns a formatted string describing the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}