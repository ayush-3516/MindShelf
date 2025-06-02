package repository

import (
	"context"

	"mindshelf/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LinkRepository defines operations for the Link entity
type LinkRepository interface {
	Create(ctx context.Context, link *model.Link) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.Link, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*model.Link, error)
	Update(ctx context.Context, link *model.Link) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Search(ctx context.Context, userID primitive.ObjectID, query string) ([]*model.Link, error)
}

// MongoLinkRepository implements LinkRepository with MongoDB
type MongoLinkRepository struct {
	repo *MongoRepository
}

// NewMongoLinkRepository creates a new MongoDB link repository
func NewMongoLinkRepository(repo *MongoRepository) *MongoLinkRepository {
	return &MongoLinkRepository{
		repo: repo,
	}
}

// Collection returns the links collection
func (r *MongoLinkRepository) Collection() *mongo.Collection {
	return r.repo.GetCollection("links")
}

// Create creates a new link in the database
func (r *MongoLinkRepository) Create(ctx context.Context, link *model.Link) error {
	result, err := r.Collection().InsertOne(ctx, link)
	if err != nil {
		return err
	}

	// Set the ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		link.ID = oid
	}

	return nil
}

// FindByID finds a link by ID
func (r *MongoLinkRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.Link, error) {
	var link model.Link
	err := r.Collection().FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Link not found
		}
		return nil, err
	}
	return &link, nil
}

// FindByUserID finds all links for a specific user
func (r *MongoLinkRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]*model.Link, error) {
	cursor, err := r.Collection().Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var links []*model.Link
	for cursor.Next(ctx) {
		var link model.Link
		if err := cursor.Decode(&link); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

// Update updates a link in the database
func (r *MongoLinkRepository) Update(ctx context.Context, link *model.Link) error {
	_, err := r.Collection().ReplaceOne(ctx, bson.M{"_id": link.ID}, link)
	return err
}

// Delete deletes a link from the database
func (r *MongoLinkRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection().DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// Search searches for links based on a query string
func (r *MongoLinkRepository) Search(ctx context.Context, userID primitive.ObjectID, query string) ([]*model.Link, error) {
	filter := bson.M{
		"user_id": userID,
		"$or": []bson.M{
			{"title": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
			{"url": bson.M{"$regex": query, "$options": "i"}},
			{"tags": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := r.Collection().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var links []*model.Link
	for cursor.Next(ctx) {
		var link model.Link
		if err := cursor.Decode(&link); err != nil {
			return nil, err
		}
		links = append(links, &link)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
