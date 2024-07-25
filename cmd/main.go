package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/Reaper1994/go-package-master/internal/config"
    v1 "github.com/Reaper1994/go-package-master/internal/handlers/v1"
    "github.com/Reaper1994/go-package-master/internal/middleware"
    "github.com/Reaper1994/go-package-master/internal/services"
)

func main() {
    cfg, err := config.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize pack calculators for each version
        calculatorV1 := services.PackCalculatorV1{Packs: cfg.Packs}

        // Handlers for each version
        handlerV1 := &v1.CalculateHandlerV1{Calculator: calculatorV1}

        // Set up routes with middleware
        http.Handle("/api/v1/calculate", middleware.LoggingMiddleware(middleware.RecoveryMiddleware(handlerV1)))

        fmt.Println("PackMaster server is running on port 8080")
        log.Fatal(http.ListenAndServe(":8080", nil))
}
