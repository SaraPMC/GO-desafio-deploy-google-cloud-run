package usecase

import (
	"fmt"

	"github.com/SaraPMC/GO-desafio-deploy-google-cloud-run/internal/entity"
	"github.com/SaraPMC/GO-desafio-deploy-google-cloud-run/internal/infra/service"
)

// GetWeatherByZipCodeUseCase is the use case for getting weather by zipcode
type GetWeatherByZipCodeUseCase struct {
	viaCEPService  *service.ViaCEPService
	weatherService *service.WeatherService
}

// NewGetWeatherByZipCodeUseCase creates a new instance of GetWeatherByZipCodeUseCase
func NewGetWeatherByZipCodeUseCase(
	viaCEPService *service.ViaCEPService,
	weatherService *service.WeatherService,
) *GetWeatherByZipCodeUseCase {
	return &GetWeatherByZipCodeUseCase{
		viaCEPService:  viaCEPService,
		weatherService: weatherService,
	}
}

// Execute executes the use case to get weather by zipcode
func (u *GetWeatherByZipCodeUseCase) Execute(zipCode string) (*entity.Weather, error) {
	// Validate and get location from ViaCEP
	location, err := u.viaCEPService.GetLocation(zipCode)
	if err != nil {
		// Return the error as is (it already has the proper message)
		return nil, err
	}

	// Get weather for the city
	weather, err := u.weatherService.GetWeather(location.City)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %w", err)
	}

	return weather, nil
}
