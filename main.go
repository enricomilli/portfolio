package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/enricomilli/portfolio/site"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	r := chi.NewRouter()

	err := setupEnv()
	if err != nil {
		return fmt.Errorf("failed to initialize environment: %w", err)
	}

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
}

func setupFileServer(r *chi.Mux) {
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
}

func setupPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func setupEnv() error {
	if os.Getenv("K_SERVICE") == "" {
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf("failed to load .env file: %w", err)
		}
	}
	return nil
}
