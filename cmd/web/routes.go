package main

import (
	"bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(LoadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
