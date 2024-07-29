package main

import (
	"flag"
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

const (
	Port = 8080
)

var enableTreblle = flag.Bool("enable-treblle", true, "Enable Treblle integration")

func initializeMiddleware(enableTreblleMiddleware bool) http.Handler {
	// Initialize pack calculators for each version
	calculatorV1 := services.PackCalculatorV1{}

	// Handlers for each version.
	handlerV1 := &v1.CalculateHandlerV1{Calculator: calculatorV1}

	var treblleHandler http.Handler

	if enableTreblleMiddleware {
		treblleAPIKey := os.Getenv("TREBLLE_API_KEY")
		treblleProjectID := os.Getenv("TREBLLE_PROJECT_ID")

		treblleHandler = middleware.TreblleMiddleware(treblleAPIKey, treblleProjectID, handlerV1)
	}

	// Apply custom middleware
	finalHandler := middleware.LoggingMiddleware(middleware.RecoveryMiddleware(treblleHandler))

	return finalHandler
}

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Convert cfg.Packs to []models.Pack
	var packs []models.Pack
	for _, p := range cfg.Packs {
		packs = append(packs, models.Pack{Size: p.Size})
	}

	http.Handle("/api/v1/calculate", initializeMiddleware(*enableTreblle))

	fmt.Printf("PackMaster server is running on port %d\n", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
