package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetWeatherData retrieves weather data for the specified latitude and longitude.
func GetWeatherData(lat, lon float64) (WeatherData, error) {
	var data WeatherData

	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current=temperature_2m,weather_code&daily=temperature_2m_max,temperature_2m_min&forecast_days=1",
		strconv.FormatFloat(lat, 'f', -1, 64),
		strconv.FormatFloat(lon, 'f', -1, 64),
	)

	resp, err := http.Get(weatherURL)
	if err != nil {
		return data, fmt.Errorf("error fetching weather data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("weather API returned status %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error decoding weather response: %w", err)
	}

	return data, nil
}
