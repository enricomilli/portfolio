package home

import "github.com/go-chi/chi/v5"

func SetupRoutes(homeRouter chi.Router) {

	homeRouter.Get("/", HandleIndexPage)
	homeRouter.Get("/random-number", HandleRandNum)

	// form state component
	homeRouter.Put("/form-state", HandleFormStatePut)

}
