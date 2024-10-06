package home

import "github.com/go-chi/chi/v5"

func SetupRoutes(indexRouter chi.Router) {

	indexRouter.Get("/", HandleIndex)

	// pages
	indexRouter.Get("/home", HandleHomePage)
	indexRouter.Post("/search", HandleSearchPage)

	// component routes
	indexRouter.Get("/random-number", HandleRandNum)
	indexRouter.Put("/form-state", HandleFormStatePut)

}
