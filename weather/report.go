package weather

import (
	"fmt"
	"math"
)

// GetWeatherReport fetches and prints the weather report for the given location.
// It expects the path to the weather codes JSON file.
func GetWeatherReport(location string) error {
	fmt.Println("...Fetching weather report for:", location)

	geo, err := GetGeoCoordinates(location)
	if err != nil {
		fmt.Println("Error from GetGeoCoordinates", err)
		return err
	}

	weatherData, err := GetWeatherData(geo.Latitude, geo.Longitude)
	if err != nil {
		fmt.Println("Error from GetWeatherData", err)
		return err
	}

	if err := LoadWeatherCodes(); err != nil {
		fmt.Println("Error from LoadWeatherCodes", err)
		return err
	}

	description, err := GetWeatherCodeDescription(weatherData.Current.WeatherCode)
	if err != nil {
		fmt.Println("Error from GetWeatherCodeDescription", err)
		return err
	}

	fmt.Printf("The weather in %s is: %s\n", location, description)
	fmt.Printf("The current temperature: %.0f°C\n", math.Round(weatherData.Current.Temperature2m))

	if len(weatherData.Daily.Temperature2mMax) > 0 {
		fmt.Printf("The daily max temperature: %.0f°C\n", math.Round(weatherData.Daily.Temperature2mMax[0]))
	}
	if len(weatherData.Daily.Temperature2mMin) > 0 {
		fmt.Printf("The daily min temperature: %.0f°C\n", math.Round(weatherData.Daily.Temperature2mMin[0]))
	}

	return nil
}
