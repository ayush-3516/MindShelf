package service

import (
	"context"
	"errors"
	"time"

	"mindshelf/internal/model"
	"mindshelf/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkService provides link-related functionality
type LinkService struct {
	linkRepo repository.LinkRepository
}

// NewLinkService creates a new link service
func NewLinkService(linkRepo repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepo: linkRepo,
	}
}

// CreateLink creates a new link
func (s *LinkService) CreateLink(ctx context.Context, userID primitive.ObjectID, req *model.CreateLinkRequest) (*model.Link, error) {
	now := time.Now()
	link := &model.Link{
		UserID:      userID,
		URL:         req.URL,
		Title:       req.Title,
		Description: req.Description,
		Tags:        req.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.linkRepo.Create(ctx, link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

// GetLink gets a link by ID
func (s *LinkService) GetLink(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (*model.Link, error) {
	link, err := s.linkRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if link == nil {
		return nil, errors.New("link not found")
	}

	// Ensure the link belongs to the user
	if link.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return link, nil
}

// GetUserLinks gets all links for a user
func (s *LinkService) GetUserLinks(ctx context.Context, userID primitive.ObjectID) ([]*model.Link, error) {
	return s.linkRepo.FindByUserID(ctx, userID)
}

// UpdateLink updates a link
func (s *LinkService) UpdateLink(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, req *model.UpdateLinkRequest) (*model.Link, error) {
	link, err := s.linkRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if link == nil {
		return nil, errors.New("link not found")
	}

	// Ensure the link belongs to the user
	if link.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	// Update link fields
	link.URL = req.URL
	link.Title = req.Title
	link.Description = req.Description
	link.Tags = req.Tags
	link.UpdatedAt = time.Now()

	err = s.linkRepo.Update(ctx, link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

// DeleteLink deletes a link
func (s *LinkService) DeleteLink(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	link, err := s.linkRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if link == nil {
		return errors.New("link not found")
	}

	// Ensure the link belongs to the user
	if link.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.linkRepo.Delete(ctx, id)
}

// SearchLinks searches for links
func (s *LinkService) SearchLinks(ctx context.Context, userID primitive.ObjectID, query string) ([]*model.Link, error) {
	return s.linkRepo.Search(ctx, userID, query)
}
