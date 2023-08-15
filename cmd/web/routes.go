package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luigieai/go-booking-website/pkg/config"
	"github.com/luigieai/go-booking-website/pkg/handlers"
)

func Routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	// Temporary :)
	tempHtmlServer := http.FileServer(http.Dir("./temphtml/"))
	mux.Handle("/html/*", http.StripPrefix("/html", tempHtmlServer))
	return mux
}
