package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // Password is not returned in JSON
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// LoginRequest represents user login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest represents user registration data
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWTClaims represents JWT claims
type JWTClaims struct {
	UserID string `json:"user_id"`
}
