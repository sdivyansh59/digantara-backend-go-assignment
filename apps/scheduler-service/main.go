package main

import (
	"os"
	"path/filepath"
	_ "time/tzdata" // ensure we always have the timezone information included

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/sdivyansh59/digantara-backend-golang-assignment/app"
)

func main() {
	// Load .env file first before any initialization
	loadEnvironmentVariables()

	// Wire handles all initialization including logger setup
	application, err := app.InitializeApp()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize application")
	}

	if err := application.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run application")
	}
}

func loadEnvironmentVariables() {
	// Load .env file from current directory
	if err := godotenv.Load(".env"); err != nil {
		// Try to load from the executable's directory as fallback
		if ex, err := os.Executable(); err == nil {
			exPath := filepath.Dir(ex)
			_ = godotenv.Load(filepath.Join(exPath, ".env"))
		}
		// Not a fatal error - can use system environment variables
		log.Warn().Msg("No .env file found, using system environment variables")
	} else {
		log.Info().Msg(".env file loaded successfully")
	}
}
