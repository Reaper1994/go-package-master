package services_test

import (
	"log"
	"testing"

	"github.com/Reaper1994/go-package-master/internal/config"
	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
	"github.com/Reaper1994/go-package-master/internal/transformers"
	"github.com/stretchr/testify/assert"
)

var calculator services.PackCalculatorV1

func init() {
	cfg, err := config.LoadConfig("../cmd/config.json") // adjust the path as needed
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var packs []models.Pack
	for _, p := range cfg.Packs {
		packs = append(packs, models.Pack{Size: p.Size})
	}

	calculator = services.PackCalculatorV1{
		Packs: packs,
	}
}

func TestPackCalculatorV1_CalculatePacks_1(t *testing.T) {
	order := models.Order{Items: 1}
	expectedJSON := `{"250":1}`

	packs := calculator.CalculatePacks(order)
	result := transformers.FormatPacks(packs)

	assert.JSONEq(t, expectedJSON, result, "Expected packs %s, but got %s", expectedJSON, result)
}

func TestPackCalculatorV1_CalculatePacks_250(t *testing.T) {
	order := models.Order{Items: 250}
	expectedJSON := `{"250":1}`

	packs := calculator.CalculatePacks(order)
	result := transformers.FormatPacks(packs)

	assert.JSONEq(t, expectedJSON, result, "Expected packs %s, but got %s", expectedJSON, result)
}

func TestPackCalculatorV1_CalculatePacks_251(t *testing.T) {
	order := models.Order{Items: 251}
	expectedJSON := `{"500":1}`

	packs := calculator.CalculatePacks(order)
	result := transformers.FormatPacks(packs)

	assert.JSONEq(t, expectedJSON, result, "Expected packs %s, but got %s", expectedJSON, result)
}

func TestPackCalculatorV1_CalculatePacks_501(t *testing.T) {
	order := models.Order{Items: 501}
	expectedJSON := `{"500":1,"250":1}`

	packs := calculator.CalculatePacks(order)
	result := transformers.FormatPacks(packs)

	assert.JSONEq(t, expectedJSON, result, "Expected packs %s, but got %s", expectedJSON, result)
}

func TestPackCalculatorV1_CalculatePacks_12001(t *testing.T) {
	order := models.Order{Items: 12001}
	expectedJSON := `{"5000":2,"2000":1,"250":1}`

	packs := calculator.CalculatePacks(order)
	result := transformers.FormatPacks(packs)

	assert.JSONEq(t, expectedJSON, result, "Expected packs %s, but got %s", expectedJSON, result)
}

func TestPackCalculatorV1_CalculatePacks_Invalid(t *testing.T) {
	// Testing with order amount 0
	orderZero := models.Order{Items: 0}
	expectedJSONZero := `{}`

	packsZero := calculator.CalculatePacks(orderZero)
	resultZero := transformers.FormatPacks(packsZero)
	assert.JSONEq(t, expectedJSONZero, resultZero, "Expected packs %s, but got %s for order amount 0", expectedJSONZero, resultZero)

	// Testing with negative order amount
	orderNegative := models.Order{Items: -1}
	expectedJSONNegative := `{}`

	packsNegative := calculator.CalculatePacks(orderNegative)
	resultNegative := transformers.FormatPacks(packsNegative)
	assert.JSONEq(t, expectedJSONNegative, resultNegative, "Expected packs %s, but got %s for negative order amount", expectedJSONNegative, resultNegative)
}
