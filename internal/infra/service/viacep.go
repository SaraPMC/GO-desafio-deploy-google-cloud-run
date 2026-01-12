package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/sarapmc/weather-app/internal/entity"
)

// ViaCEPService is responsible for fetching location data from ViaCEP API
type ViaCEPService struct {
	httpClient *http.Client
	baseURL    string
}

// NewViaCEPService creates a new instance of ViaCEPService
func NewViaCEPService() *ViaCEPService {
	return &ViaCEPService{
		httpClient: &http.Client{},
		baseURL:    "https://viacep.com.br/ws",
	}
}

// ValidateZipCode validates if the zipcode has the correct format (8 digits)
func (v *ViaCEPService) ValidateZipCode(zipCode string) error {
	// Remove any non-digit characters
	cleaned := regexp.MustCompile(`\D`).ReplaceAllString(zipCode, "")

	// Check if it has exactly 8 digits
	if len(cleaned) != 8 {
		return fmt.Errorf("invalid zipcode")
	}

	return nil
}

// GetLocation fetches location data from ViaCEP API
func (v *ViaCEPService) GetLocation(zipCode string) (*entity.Location, error) {
	// Validate zipcode format
	if err := v.ValidateZipCode(zipCode); err != nil {
		return nil, err
	}

	// Clean zipcode (remove non-digits)
	cleaned := regexp.MustCompile(`\D`).ReplaceAllString(zipCode, "")

	// Format URL with zipcode
	url := fmt.Sprintf("%s/%s/json", v.baseURL, cleaned)

	// Make HTTP request
	resp, err := v.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}

	// Parse JSON response
	var location entity.Location
	if err := json.Unmarshal(body, &location); err != nil {
		return nil, fmt.Errorf("can not find zipcode")
	}

	// Check if location was found
	if location.Error || strings.TrimSpace(location.City) == "" {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return &location, nil
}
