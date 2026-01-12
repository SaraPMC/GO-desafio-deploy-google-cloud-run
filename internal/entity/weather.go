package entity

// Weather represents the weather information with temperature in different units
type Weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

// Location represents the location information from ViaCEP API
type Location struct {
	CEP        string `json:"cep"`
	City       string `json:"localidade"`
	State      string `json:"uf"`
	District   string `json:"bairro"`
	Street     string `json:"logradouro"`
	Number     string `json:"numero"`
	Complement string `json:"complemento"`
	Error      bool   `json:"erro,omitempty"`
}

// WeatherResponse represents the complete response with weather data
type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
