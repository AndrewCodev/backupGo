//Package weather: Permite a los desarrolladores obtener datos meteorológicos en tiempo real y pronósticos, como temperatura, condiciones, viento y humedad, directamente desde sus aplicaciones.
package weather

var (
    //CurrentCondition indica la candición actual del clima.
	CurrentCondition string

    //CurrentLocation indica la locación actual.
	CurrentLocation  string
)

//Forecast nos dice de manera ordenada la locación y la condición del clima actual.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
