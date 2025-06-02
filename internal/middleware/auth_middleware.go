package middleware

import (
	"context"
	"net/http"
	"strings"

	"mindshelf/internal/service"
)

// UserIDKey is the context key for the user ID
type UserIDKey string

const ContextUserIDKey UserIDKey = "user_id"

// AuthMiddleware provides authentication middleware
type AuthMiddleware struct {
	authService *service.AuthService
}

// NewAuthMiddleware creates a new authentication middleware
func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Authenticate authenticates a request
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}
		tokenString := tokenParts[1]

		// Validate the token
		userID, err := m.authService.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add the user ID to the context
		ctx := context.WithValue(r.Context(), ContextUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
