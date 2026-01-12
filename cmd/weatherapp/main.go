package main

import (
	"fmt"
	"os"

	"github.com/sarapmc/weather-app/internal/infra/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting Weather App on port %s\n", port)

	if err := web.StartServer(port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
