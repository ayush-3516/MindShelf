package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"mindshelf/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Load configuration (which includes MongoDB connection string)
	cfg := config.Load()

	// Print config details (without sensitive info)
	fmt.Println("Connecting to database:", cfg.MongoDatabase)
	fmt.Println("MongoDB URI:", maskConnectionString(cfg.MongoURI))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Ping the database to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB Atlas!")

	// List available databases
	databaseNames, err := client.ListDatabaseNames(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}

	fmt.Println("Available databases:")
	for _, name := range databaseNames {
		fmt.Println("-", name)
	}
}

// maskConnectionString hides the username and password in the connection string
func maskConnectionString(uri string) string {
	// Simple implementation that just returns a placeholder
	// This is to avoid accidentally logging sensitive credentials
	return "[CONNECTION STRING HIDDEN]"
}
