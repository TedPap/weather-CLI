package weather

// GeoDataResult represents one result from the geocoding API.
type GeoDataResult struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GeoData wraps a slice of geocoding results.
type GeoData struct {
	Results []GeoDataResult `json:"results"`
}

// WeatherData represents the weather API response.
type WeatherData struct {
	Current struct {
		Temperature2m float64 `json:"temperature_2m"`
		WeatherCode   float64 `json:"weather_code"`
	} `json:"current"`
	Daily struct {
		Temperature2mMax []float64 `json:"temperature_2m_max"`
		Temperature2mMin []float64 `json:"temperature_2m_min"`
	} `json:"daily"`
}

// WeatherCodeDescription holds the description information for a weather code.
type WeatherCodeDescription struct {
	Day struct {
		Description string `json:"description"`
	} `json:"day"`
}
