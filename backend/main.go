package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ayush-3516/mindshelf/backend/internal/db"
	"github.com/ayush-3516/mindshelf/backend/repository"
	"github.com/ayush-3516/mindshelf/backend/service"
	"github.com/ayush-3516/mindshelf/backend/api"
	"github.com/ayush-3516/mindshelf/backend/router"
)

func main() {
	// 🧪 Load DB config (hardcoded or via env vars)
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "yourpassword")
	dbname := getEnv("DB_NAME", "yourdb")
	portStr := getEnv("DB_PORT", "5432")
	port, _ := strconv.Atoi(portStr)

	// 🧵 Connect to DB
	conn, err := db.ConnectDB(host, user, password, dbname, port)
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer conn.Close()

	// 🧠 Init sqlc-generated queries
	queries := db.New(conn)

	// 🧱 Init app layers
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	// 🌐 Setup router
	r := router.SetupRouter(userHandler)

	// 🚀 Start server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// 📦 Helper: fallback env var
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

