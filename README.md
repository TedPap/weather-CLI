# Weather CLI

A simple command-line tool written in Go that fetches and displays the current weather report for a given location. It uses the Open-Meteo APIs for both geocoding and weather data, and maps weather codes to human-readable descriptions using a local JSON file.

## Features

- **Geocoding**: Converts a location name into latitude and longitude using the Open-Meteo Geocoding API.
- **Weather Data**: Retrieves current weather conditions and daily forecasts from the Open-Meteo Weather API.

## Prerequisites

- [Go](https://golang.org/) 1.18 or later
- Internet connectivity (to access the Open-Meteo APIs)

## Installation

1. **Clone the Repository**
   ```bash
    git clone https://github.com/yourusername/weather-CLI.git
    cd weather-CLI
   ```
1. Build the Application
   Run the following command to build the executable:
   ```bash
     go build -o weather-CLI .
   ```

This will produce an executable named weather-CLI in the project directory.

## Usage

Run the CLI tool with the required -location flag

```bash
  ./weather-CLI -location="New York"
```

- -location: (required) The name of the location for which you want the weather report.

### Example

```bash
  ./weather-CLI -location="Los Angeles"
```

This command will fetch and display the current weather for Los Angeles.
Configuration

The tool leverages two external APIs:

    Geocoding API: Open-Meteo Geocoding API for converting location names into coordinates.
    Weather API: Open-Meteo Weather API for fetching current weather and daily forecast data.

Make sure these services are accessible from your network.
