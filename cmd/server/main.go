package main

import (
	"log"
	"net/http"

	"github.com/MalcolmFuchs/Investment-Calculator-Planner/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	address := ":8080"
	r := chi.NewRouter() // Korrekte Initialisierung

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Root-Handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome!"))
	})

	// API-Routen
	r.Route("/api", func(r chi.Router) {
		r.Post("/calculate", api.ErrorHandler(api.InvestmentRequestHandler))
	})

	log.Printf("Server running on %s", address)
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
