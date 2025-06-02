package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mindshelf/internal/middleware"
	"mindshelf/internal/model"
	"mindshelf/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LinkHandler handles link-related requests
type LinkHandler struct {
	linkService *service.LinkService
}

// NewLinkHandler creates a new link handler
func NewLinkHandler(linkService *service.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

// CreateLink handles link creation
func (h *LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	var req model.CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate request
	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	link, err := h.linkService.CreateLink(r.Context(), userID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(link)
}

// GetLink handles getting a link
func (h *LinkHandler) GetLink(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	link, err := h.linkService.GetLink(r.Context(), id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

// GetLinks handles getting all links for a user
func (h *LinkHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	links, err := h.linkService.GetUserLinks(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

// UpdateLink handles updating a link
func (h *LinkHandler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	var req model.UpdateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	link, err := h.linkService.UpdateLink(r.Context(), id, userID, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

// DeleteLink handles deleting a link
func (h *LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	err = h.linkService.DeleteLink(r.Context(), id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SearchLinks handles searching for links
func (h *LinkHandler) SearchLinks(w http.ResponseWriter, r *http.Request) {
	// Get query parameter
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// Get user ID from context
	userID := r.Context().Value(middleware.ContextUserIDKey).(primitive.ObjectID)

	links, err := h.linkService.SearchLinks(r.Context(), userID, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}
