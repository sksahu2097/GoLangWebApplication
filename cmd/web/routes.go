package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sksahu2097/go-project/pkg/config"
	"github.com/sksahu2097/go-project/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
