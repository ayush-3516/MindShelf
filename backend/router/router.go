package router

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/ayush-3516/mindshelf/backend/api"
)

func SetupRouter(handler *api.UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/user", handler.GetUserByEmail)
	r.Post("/register", handler.CreateUser)

	return r
}

