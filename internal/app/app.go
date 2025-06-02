package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"mindshelf/internal/config"
	"mindshelf/internal/handler"
	appMiddleware "mindshelf/internal/middleware"
	"mindshelf/internal/repository"
	"mindshelf/internal/service"
)

// App represents the application
type App struct {
	config      *config.Config
	router      *chi.Mux
	mongoRepo   *repository.MongoRepository
	authService *service.AuthService
	linkService *service.LinkService
}

// New creates a new application
func New() (*App, error) {
	// Load configuration
	cfg := config.Load()

	// Create MongoDB repository
	mongoRepo, err := repository.NewMongoRepository(cfg)
	if err != nil {
		return nil, err
	}

	// Create repositories
	userRepo := repository.NewMongoUserRepository(mongoRepo)
	linkRepo := repository.NewMongoLinkRepository(mongoRepo)

	// Create services
	authService := service.NewAuthService(userRepo, cfg)
	linkService := service.NewLinkService(linkRepo)

	// Create handlers
	authHandler := handler.NewAuthHandler(authService)
	linkHandler := handler.NewLinkHandler(linkService)

	// Create auth middleware
	authMiddleware := appMiddleware.NewAuthMiddleware(authService)

	// Create router
	router := chi.NewRouter()

	// Add middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Public routes
	router.Route("/api", func(r chi.Router) {
		// Auth routes
		r.Post("/auth/register", authHandler.Register)
		r.Post("/auth/login", authHandler.Login)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(authMiddleware.Authenticate)

			// Link routes
			r.Post("/links", linkHandler.CreateLink)
			r.Get("/links", linkHandler.GetLinks)
			r.Get("/links/search", linkHandler.SearchLinks)
			r.Get("/links/{id}", linkHandler.GetLink)
			r.Put("/links/{id}", linkHandler.UpdateLink)
			r.Delete("/links/{id}", linkHandler.DeleteLink)
		})
	})

	return &App{
		config:      cfg,
		router:      router,
		mongoRepo:   mongoRepo,
		authService: authService,
		linkService: linkService,
	}, nil
}

// Start starts the application
func (a *App) Start() error {
	// Create server
	server := &http.Server{
		Addr:    ":" + a.config.ServerPort,
		Handler: a.router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s\n", a.config.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v\n", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v\n", err)
	}

	// Close MongoDB connection
	if err := a.mongoRepo.Close(ctx); err != nil {
		log.Fatalf("Error closing MongoDB connection: %v\n", err)
	}

	log.Println("Server gracefully stopped")
	return nil
}
