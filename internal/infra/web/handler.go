package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sarapmc/weather-app/internal/usecase"
)

// WeatherHandler handles HTTP requests for weather
type WeatherHandler struct {
	getWeatherByZipCodeUseCase *usecase.GetWeatherByZipCodeUseCase
}

// NewWeatherHandler creates a new instance of WeatherHandler
func NewWeatherHandler(getWeatherByZipCodeUseCase *usecase.GetWeatherByZipCodeUseCase) *WeatherHandler {
	return &WeatherHandler{
		getWeatherByZipCodeUseCase: getWeatherByZipCodeUseCase,
	}
}

// GetWeatherByZipCode handles GET requests for weather by zipcode
func (h *WeatherHandler) GetWeatherByZipCode(w http.ResponseWriter, r *http.Request) {
	// Get zipcode from query parameter
	zipCode := r.URL.Query().Get("cep")

	// Check if zipcode is provided
	if strings.TrimSpace(zipCode) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
		return
	}

	// Execute use case
	weather, err := h.getWeatherByZipCodeUseCase.Execute(zipCode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		// Check error message to determine status code
		errMsg := err.Error()
		if strings.Contains(errMsg, "invalid zipcode") {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
		} else if strings.Contains(errMsg, "can not find zipcode") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "can not find zipcode"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal server error"})
		}
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weather)
}

// HealthCheck handles health check requests
func (h *WeatherHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
