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

	// Try to access the database and a collection
	database := client.Database(cfg.MongoDatabase)
	collection := database.Collection("test_collection")

	fmt.Printf("Successfully accessed database '%s' and collection 'test_collection'\n", cfg.MongoDatabase)

	// Optional: Insert a test document to verify write permissions
	testDoc := map[string]interface{}{
		"test_field": "Hello MongoDB Atlas!",
		"timestamp":  time.Now(),
	}

	insertResult, err := collection.InsertOne(ctx, testDoc)
	if err != nil {
		log.Printf("Warning: Could not insert test document: %v", err)
		fmt.Println("Your connection is working but you might not have write permissions.")
	} else {
		fmt.Printf("Successfully inserted test document with ID: %v\n", insertResult.InsertedID)

		// Clean up - delete the test document
		_, err = collection.DeleteOne(ctx, map[string]interface{}{"_id": insertResult.InsertedID})
		if err != nil {
			log.Printf("Warning: Could not delete test document: %v", err)
		} else {
			fmt.Println("Successfully deleted test document")
		}
	}
}

// maskConnectionString hides the username and password in the connection string
func maskConnectionString(uri string) string {
	// Simple implementation that just returns a placeholder
	// This is to avoid accidentally logging sensitive credentials
	return "[CONNECTION STRING HIDDEN]"
}
