//Package weather provides us current condition of the location.
package weather

//CurrentCondition stores weather details in string.
var CurrentCondition string

//CurrentLocation stores location details in string.
var CurrentLocation string

//Forecast is a funtion which returns location and condition of location in sting.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
