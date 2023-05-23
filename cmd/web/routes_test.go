package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/markhorn-dev/go-bookings-app/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//
	default:
		t.Error(fmt.Printf("type is %T and not http.Handler", v))
	}
}
