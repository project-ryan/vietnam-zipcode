package main

import (
	"log"

	"github.com/quanghia24/vietnam-zipcode/api"
	"github.com/quanghia24/vietnam-zipcode/config"
	"github.com/quanghia24/vietnam-zipcode/service"
)

func main() {
	// Load configuration
	config := config.LoadConfig()

	// Initialize service
	service := service.NewMemoryZipcodeService()

	// Load data
	if err := service.LoadData(config.DataFile); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	log.Printf("Loaded zipcode data with %d entries", service.GetDataSize())

	// Create and start server
	server := api.NewServer(service, config)

	log.Printf("Starting Vietnam Zipcode API server on port %s", config.Port)
	if err := server.Start(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
