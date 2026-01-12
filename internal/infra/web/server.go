package web

import (
	"net/http"

	"github.com/SaraPMC/GO-desafio-deploy-google-cloud-run/internal/infra/service"
	"github.com/SaraPMC/GO-desafio-deploy-google-cloud-run/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// Router creates and configures the HTTP router
func Router() http.Handler {
	// Initialize services
	viaCEPService := service.NewViaCEPService()
	weatherService := service.NewWeatherService()

	// Initialize use cases
	getWeatherByZipCodeUseCase := usecase.NewGetWeatherByZipCodeUseCase(viaCEPService, weatherService)

	// Initialize handlers
	weatherHandler := NewWeatherHandler(getWeatherByZipCodeUseCase)

	// Create router
	r := chi.NewRouter()

	// Add CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	r.Get("/", weatherHandler.HealthCheck)
	r.Get("/weather", weatherHandler.GetWeatherByZipCode)

	return r
}

// StartServer starts the HTTP server
func StartServer(port string) error {
	if port == "" {
		port = "8080"
	}

	router := Router()

	return http.ListenAndServe(":"+port, router)
}
