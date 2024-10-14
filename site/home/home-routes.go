package home

import "github.com/go-chi/chi/v5"

func SetupRoutes(indexRouter chi.Router) {

	indexRouter.Get("/", HandleIndex)

	// pages
	indexRouter.Get("/home", HandleHomePage)
	indexRouter.Post("/search", HandleSearchPage)

	// component routes
	indexRouter.Put("/form-state", HandleFormStatePut)

	indexRouter.Route("/random-number", RandomNumberRoutes)

}

func RandomNumberRoutes(randRouter chi.Router) {

	// TODO: add range as a state in the client
	// returns the div with the random number
	randRouter.Post("/new", HandleRandNum)

	randRouter.Get("/edit", HandleRenderEditForm)

	randRouter.Put("/edit", HandleUpdatedRange)

	// randNumRouter.Get("/edit", ReturnEditRangeForm)
	// randNumRouter.Put("/edit", HandleUpdateRange)
}
