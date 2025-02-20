package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var weatherCodes map[string]WeatherCodeDescription

// LoadWeatherCodes loads weather code descriptions from the given JSON file.
func LoadWeatherCodes() error {
	filename := "./weatherCodes.json"
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading weather codes file: %w", err)
	}

	weatherCodes = make(map[string]WeatherCodeDescription)
	if err := json.Unmarshal(content, &weatherCodes); err != nil {
		return fmt.Errorf("error unmarshalling weather codes: %w", err)
	}

	return nil
}

// GetWeatherCodeDescription returns the description for the given weather code.
func GetWeatherCodeDescription(code float64) (string, error) {
	codeStr := strconv.FormatFloat(code, 'f', -1, 64)
	wc, ok := weatherCodes[codeStr]
	if !ok {
		return "", fmt.Errorf("unsupported weather code: %s", codeStr)
	}

	if wc.Day.Description == "" {
		return "", fmt.Errorf("weather code %s has no description", codeStr)
	}

	return wc.Day.Description, nil
}
