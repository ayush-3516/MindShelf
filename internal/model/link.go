package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Link represents a saved link in the system
type Link struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	URL         string             `bson:"url" json:"url"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// CreateLinkRequest represents the data needed to create a new link
type CreateLinkRequest struct {
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// UpdateLinkRequest represents the data needed to update a link
type UpdateLinkRequest struct {
	URL         string   `json:"url"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
