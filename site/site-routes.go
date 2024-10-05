package site

import (
	"github.com/enricomilli/portfolio/site/home"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {

	// Public Routes
	r.Route("/", home.SetupRoutes)
}
