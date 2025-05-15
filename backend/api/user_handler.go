package api

import (
	"encoding/json"
	"net/http"

	"github.com/ayush-3516/mindshelf/backend/internal/db"
	"github.com/ayush-3516/mindshelf/backend/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// GET /user?email=someone@example.com
func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// POST /register
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	arg := db.CreateUserParams{
		Email:        input.Email,
		PasswordHash: input.Password, // ideally, hash this first
		Name:         toText(input.Name),
	}

	user, err := h.service.CreateUser(r.Context(), arg)
	if err != nil {
		http.Error(w, "could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Helper to wrap string into pgtype.Text
func toText(s string) db.PgText {
	return db.PgText{String: s, Valid: true}
}

