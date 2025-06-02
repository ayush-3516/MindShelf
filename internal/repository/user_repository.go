package repository

import (
	"context"

	"mindshelf/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository defines operations for the User entity
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

// MongoUserRepository implements UserRepository with MongoDB
type MongoUserRepository struct {
	repo *MongoRepository
}

// NewMongoUserRepository creates a new MongoDB user repository
func NewMongoUserRepository(repo *MongoRepository) *MongoUserRepository {
	return &MongoUserRepository{
		repo: repo,
	}
}

// Collection returns the users collection
func (r *MongoUserRepository) Collection() *mongo.Collection {
	return r.repo.GetCollection("users")
}

// Create creates a new user in the database
func (r *MongoUserRepository) Create(ctx context.Context, user *model.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert user into the database
	result, err := r.Collection().InsertOne(ctx, user)
	if err != nil {
		return err
	}

	// Set the ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return nil
}

// FindByID finds a user by ID
func (r *MongoUserRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*model.User, error) {
	var user model.User
	err := r.Collection().FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds a user by email
func (r *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.Collection().FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}
