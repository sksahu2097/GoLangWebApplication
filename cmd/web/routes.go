package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/sksahu2097/go-project/pkg/config"
	"github.com/sksahu2097/go-project/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}
