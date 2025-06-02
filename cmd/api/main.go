package main

import (
	"log"

	"mindshelf/internal/app"
)

func main() {
	// Create app
	app, err := app.New()
	if err != nil {
		log.Fatalf("Error creating app: %v", err)
	}

	// Start app
	if err := app.Start(); err != nil {
		log.Fatalf("Error starting app: %v", err)
	}
}
