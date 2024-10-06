package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/enricomilli/portfolio/site"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

//go:embed public
var FS embed.FS

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	r := chi.NewRouter()

	// err := setupEnv()
	// if err != nil {
	// 	return fmt.Errorf("failed to initialize environment: %w", err)
	// }

	setupMiddleware(r)

	setupFileServer(r)

	site.SetupRoutes(r)

	port := setupPort()

	fmt.Printf("Starting server on port %s\n", port)

	return http.ListenAndServe(port, r)
}

func setupMiddleware(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(httprate.Limit(
		1000,
		30*time.Second,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error": "Rate limited. Please slow down."}`, http.StatusTooManyRequests)
		}),
	))
}

func setupFileServer(r *chi.Mux) {
	fileServer := http.FileServer(http.Dir("./static"))
	// serving static folder
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	// serving public folder
	r.Handle("/public/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
}

func setupPort() string {
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	port := "8080"
	return ":" + port
}

// func setupEnv() error {
// 	if os.Getenv("K_SERVICE") == "" {
// 		if err := godotenv.Load(); err != nil {
// 			return fmt.Errorf("failed to load .env file: %w", err)
// 		}
// 	}
// 	return nil
// }
