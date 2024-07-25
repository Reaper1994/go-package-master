package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
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

	packs := h.Calculator.CalculatePacks(order)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(packs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}