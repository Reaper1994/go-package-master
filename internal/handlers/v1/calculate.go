package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
	"github.com/Reaper1994/go-package-master/internal/transformers"
)

// CalculateHandlerV1 handles requests to calculate pack combinations for API v1.
type CalculateHandlerV1 struct {
	Calculator services.PackCalculatorV1
}

// ServeHTTP processes the request and responds with the calculated packs.
func (h *CalculateHandlerV1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check for invalid order items
	if order.Items <= 0 {
		http.Error(w, "Did you forget to order something?", http.StatusBadRequest)
		return
	}

	// Check Accept header
	acceptHeader := r.Header.Get("Accept")
	if !strings.Contains(acceptHeader, "application/json") {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	packs := h.Calculator.CalculatePacks(order)
	formattedPacks := transformers.FormatPacks(packs) // formats the packs to give a count of each pack

	// Ensure formattedPacks is a map
	var jsonPacks map[string]int
	if err := json.Unmarshal([]byte(formattedPacks), &jsonPacks); err != nil {
		http.Error(w, "Failed to parse formatted packs", http.StatusInternalServerError)
		return
	}

	// Set security headers
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'; object-src 'none'")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("Allow", "POST")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(jsonPacks); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
