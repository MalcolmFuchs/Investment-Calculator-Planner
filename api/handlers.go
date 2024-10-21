package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MalcolmFuchs/Investment-Calculator-Planner/internal/investment"
)

// HandlerFunc ist ein benutzerdefinierter Handler-Typ, der einen Fehler zurückgibt
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler ist eine Middleware, die Fehler von Handlern verarbeitet
func ErrorHandler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			log.Println("Handler error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// InvestmentRequestHandler verarbeitet Investmentberechnungsanfragen
func InvestmentRequestHandler(w http.ResponseWriter, r *http.Request) error {
	var req investment.InvestmentRequest

	// Dekodiere die JSON-Anfrage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding request:", err)
		return err
	}

	// Instanziiere den Calculator
	calculator := &investment.Calculator{}

	// Führe die Berechnung durch
	response, err := calculator.CalculateFutureValue(req)
	if err != nil {
		log.Println("Error calculating future value:", err)
		return err
	}

	// Kodieren der Antwort als JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response:", err)
		return err
	}

	return nil
}
