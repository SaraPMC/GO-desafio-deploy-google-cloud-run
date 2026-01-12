package service

import (
	"testing"
)

func TestValidateZipCode(t *testing.T) {
	viaCEP := NewViaCEPService()

	testCases := []struct {
		name      string
		zipCode   string
		wantError bool
	}{
		{
			name:      "Valid zipcode with 8 digits",
			zipCode:   "01310100",
			wantError: false,
		},
		{
			name:      "Valid zipcode with hyphen",
			zipCode:   "01310-100",
			wantError: false,
		},
		{
			name:      "Invalid zipcode with less than 8 digits",
			zipCode:   "123",
			wantError: true,
		},
		{
			name:      "Invalid zipcode with more than 8 digits",
			zipCode:   "01310100123",
			wantError: true,
		},
		{
			name:      "Invalid zipcode with letters",
			zipCode:   "ABCDEFGH",
			wantError: true,
		},
		{
			name:      "Empty zipcode",
			zipCode:   "",
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := viaCEP.ValidateZipCode(tc.zipCode)
			if (err != nil) != tc.wantError {
				t.Errorf("ValidateZipCode(%s) error = %v, wantError %v", tc.zipCode, err, tc.wantError)
			}
		})
	}
}

func TestGetLocation(t *testing.T) {
	viaCEP := NewViaCEPService()

	testCases := []struct {
		name      string
		zipCode   string
		wantError bool
		wantCity  string
	}{
		{
			name:      "Valid zipcode - São Paulo",
			zipCode:   "01310100",
			wantError: false,
			wantCity:  "São Paulo",
		},
		{
			name:      "Invalid zipcode format",
			zipCode:   "123",
			wantError: true,
			wantCity:  "",
		},
		{
			name:      "Nonexistent zipcode",
			zipCode:   "99999999",
			wantError: true,
			wantCity:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			location, err := viaCEP.GetLocation(tc.zipCode)
			if (err != nil) != tc.wantError {
				t.Errorf("GetLocation(%s) error = %v, wantError %v", tc.zipCode, err, tc.wantError)
				return
			}
			if !tc.wantError && location.City != tc.wantCity {
				t.Errorf("GetLocation(%s) city = %s, want %s", tc.zipCode, location.City, tc.wantCity)
			}
		})
	}
}
