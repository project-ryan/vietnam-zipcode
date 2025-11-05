package main

import (
	"log"

	"github.com/quanghia24/vietnam-zipcode/api"
	"github.com/quanghia24/vietnam-zipcode/db"
	"github.com/quanghia24/vietnam-zipcode/utils"
)

func main() {
	// Load data
	data, err := utils.LoadData("./db/data.json")
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	store := db.NewStore(data)

	server := api.NewServer(store)

	if err := server.Start("0.0.0.0:8080"); err != nil {
		log.Fatal("cannot start server", err)
	}
}
