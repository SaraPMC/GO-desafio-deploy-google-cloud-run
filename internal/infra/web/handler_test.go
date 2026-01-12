package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sarapmc/weather-app/internal/entity"
	"github.com/sarapmc/weather-app/internal/infra/service"
	"github.com/sarapmc/weather-app/internal/usecase"
)

type mockViaCEPService struct {
	shouldReturnError bool
	shouldReturnCity  bool
}

func (m *mockViaCEPService) ValidateZipCode(zipCode string) error {
	return nil
}

func (m *mockViaCEPService) GetLocation(zipCode string) (*entity.Location, error) {
	if m.shouldReturnError {
		return nil, nil
	}
	if m.shouldReturnCity {
		return &entity.Location{
			CEP:  zipCode,
			City: "São Paulo",
		}, nil
	}
	return nil, nil
}

type mockWeatherService struct {
	shouldReturnError bool
}

func (m *mockWeatherService) GetWeather(city string) (*entity.Weather, error) {
	if m.shouldReturnError {
		return nil, nil
	}
	return &entity.Weather{
		TempC: 28.5,
		TempF: 83.3,
		TempK: 301.5,
	}, nil
}

func TestHealthCheck(t *testing.T) {
	viaCEPService := service.NewViaCEPService()
	weatherService := service.NewWeatherService()
	useCase := usecase.NewGetWeatherByZipCodeUseCase(viaCEPService, weatherService)
	handler := NewWeatherHandler(useCase)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.HealthCheck).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"ok"}`
	var expectedMap, actualMap map[string]interface{}
	json.Unmarshal([]byte(expected), &expectedMap)
	json.Unmarshal(rr.Body.Bytes(), &actualMap)

	if actualMap["status"] != expectedMap["status"] {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetWeatherByZipCodeWithoutCEP(t *testing.T) {
	viaCEPService := service.NewViaCEPService()
	weatherService := service.NewWeatherService()
	useCase := usecase.NewGetWeatherByZipCodeUseCase(viaCEPService, weatherService)
	handler := NewWeatherHandler(useCase)

	req, err := http.NewRequest("GET", "/weather", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetWeatherByZipCode).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}

func TestGetWeatherByZipCodeWithInvalidCEP(t *testing.T) {
	viaCEPService := service.NewViaCEPService()
	weatherService := service.NewWeatherService()
	useCase := usecase.NewGetWeatherByZipCodeUseCase(viaCEPService, weatherService)
	handler := NewWeatherHandler(useCase)

	req, err := http.NewRequest("GET", "/weather?cep=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetWeatherByZipCode).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response["message"] != "invalid zipcode" {
		t.Errorf("handler returned unexpected message: got %s want 'invalid zipcode'", response["message"])
	}
}
