package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/SaraPMC/GO-desafio-deploy-google-cloud-run/internal/entity"
)

// WeatherAPIResponse represents the response from WeatherAPI
type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

// WeatherService is responsible for fetching weather data
type WeatherService struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// NewWeatherService creates a new instance of WeatherService
func NewWeatherService() *WeatherService {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		// Default API key for testing (you should use environment variable)
		apiKey = "test-key"
	}

	return &WeatherService{
		httpClient: &http.Client{},
		apiKey:     apiKey,
		baseURL:    "https://api.weatherapi.com/v1",
	}
}

// GetWeather fetches current weather for a given city
func (w *WeatherService) GetWeather(city string) (*entity.Weather, error) {
	// Build URL for WeatherAPI with proper encoding
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("%s/current.json?key=%s&q=%s", w.baseURL, w.apiKey, encodedCity)

	// Make HTTP request
	resp, err := w.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: unable to fetch weather data for city: %s", city)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading weather response: %w", err)
	}

	// Parse JSON response
	var apiResp WeatherAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("error parsing weather response: %w", err)
	}

	// Convert temperatures
	tempC := apiResp.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	return &entity.Weather{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
