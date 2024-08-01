package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Reaper1994/go-package-master/internal/config"
	v1 "github.com/Reaper1994/go-package-master/internal/handlers/v1"
	"github.com/Reaper1994/go-package-master/internal/middleware"
	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
)

const Port = 8080

func initializeMiddleware(handler http.Handler) http.Handler {

	treblleAPIKey := os.Getenv("TREBLLE_API_KEY")
	treblleProjectID := os.Getenv("TREBLLE_PROJECT_ID")

	//TODO: Remove this later
	fmt.Printf("TREBLLE_API_KEY %s\n", os.Getenv("TREBLLE_API_KEY"))
	fmt.Printf("TREBLLE_PROJECT_ID on port %s\n", os.Getenv("TREBLLE_PROJECT_ID"))

	handler = middleware.TreblleMiddleware(treblleAPIKey, treblleProjectID, handler)
	handler = middleware.LoggingMiddleware(handler)
	handler = middleware.RecoveryMiddleware(handler)
	handler = middleware.AuthorizationMiddleware(handler)

	return handler
}

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Convert cfg.Packs to []models.Pack
	var packs []models.Pack
	for _, p := range cfg.Packs {
		packs = append(packs, models.Pack{Size: p.Size})
	}

	// Initialize pack calculators for each version
	calculatorV1 := services.PackCalculatorV1{Packs: packs}

	// Handlers for each version
	handlerV1 := &v1.CalculateHandlerV1{Calculator: calculatorV1}

	// Set up routes with middleware
	http.Handle("/api/v1/calculate", initializeMiddleware(handlerV1))

	fmt.Printf("PackMaster server is running on port %d\n", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)) //logs every request
}
