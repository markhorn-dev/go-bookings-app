package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/markhorn-dev/go-bookings-app/pkg/config"
	"github.com/markhorn-dev/go-bookings-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/availability", handlers.Repo.Availability)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/generals-suite", handlers.Repo.GeneralsSuite)
	mux.Get("/majors-quarters", handlers.Repo.MajorsQuarters)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
