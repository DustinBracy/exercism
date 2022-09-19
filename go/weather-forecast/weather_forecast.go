// Package weather provides a weather forecast for a given city and condition.
package weather

//CurrentCondition provides the current weather condition.
var CurrentCondition string
//CurrentLocation identifies the city in which the weather applies.
var CurrentLocation string

//Forecast gives the current weather condition for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
