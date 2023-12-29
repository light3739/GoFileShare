package main

import (
	"fmt"
	"log"
	"net/http"
	"share-files/api"
	"share-files/configs"
	"share-files/db"
	"share-files/logger"
)

func main() {
	err := configs.Init()
	if err != nil {
		log.Fatalf("Error loading configurations: %s", err)
	}

	// Connect to the database
	dbConn, err := configs.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	defer dbConn.Close()

	// Run migrations
	db.Migrate(dbConn)

	// @title Share-Files
	// @version 1.0
	// @description Sosi suka
	// @host localhost:8080
	// @BasePath /

	r := api.SetupRouter()
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to start server")
	}
}
