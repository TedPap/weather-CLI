package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetGeoCoordinates returns the latitude and longitude for a given location.
func GetGeoCoordinates(location string) (GeoDataResult, error) {
	var result GeoDataResult

	// Prepare the location string for the URL.
	urlLocation := strings.ReplaceAll(location, " ", "+")
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", urlLocation)

	resp, err := http.Get(geoURL)
	if err != nil {
		return result, fmt.Errorf("error fetching geocoding data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("geocoding API returned status %s", resp.Status)
	}

	var geoData GeoData
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return result, fmt.Errorf("error decoding geocoding response: %w", err)
	}

	if len(geoData.Results) == 0 {
		return result, fmt.Errorf("no results found for location: %s", location)
	}

	return geoData.Results[0], nil
}
