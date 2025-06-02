package repository

import (
	"context"
	"time"

	"mindshelf/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoRepository provides access to MongoDB
type MongoRepository struct {
	client   *mongo.Client
	database *mongo.Database
}

// NewMongoRepository creates a new MongoDB repository
func NewMongoRepository(cfg *config.Config) (*MongoRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	// Get database
	database := client.Database(cfg.MongoDatabase)

	return &MongoRepository{
		client:   client,
		database: database,
	}, nil
}

// Close closes the MongoDB connection
func (r *MongoRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection
func (r *MongoRepository) GetCollection(name string) *mongo.Collection {
	return r.database.Collection(name)
}
