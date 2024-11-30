package main

import (
	"github/iamlego/go-web/pkg/config"
	handler "github/iamlego/go-web/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/About", http.HandlerFunc(handler.Repo.About))
	mux := chi.NewRouter()

	mux.Use(WriteToDConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/About", handler.Repo.About)

	return mux
}
