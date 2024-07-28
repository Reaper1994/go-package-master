package services_test

import (
	"log"
	"testing"

	"github.com/Reaper1994/go-package-master/internal/config"
	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
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
	expectedPacks := []models.Pack{{Size: 250}}

	packs := calculator.CalculatePacks(order)
	if !equalPacks(packs, expectedPacks) {
		t.Errorf("Expected packs %v, but got %v", expectedPacks, packs)
	} else {
		t.Log("TestPackCalculatorV1_CalculatePacks_1: Success")
	}
}

func TestPackCalculatorV1_CalculatePacks_250(t *testing.T) {
	order := models.Order{Items: 250}
	expectedPacks := []models.Pack{{Size: 250}}

	packs := calculator.CalculatePacks(order)
	if !equalPacks(packs, expectedPacks) {
		t.Errorf("Expected packs %v, but got %v", expectedPacks, packs)
	} else {
		t.Log("TestPackCalculatorV1_CalculatePacks_250: Success")
	}
}

func TestPackCalculatorV1_CalculatePacks_251(t *testing.T) {
	order := models.Order{Items: 251}
	expectedPacks := []models.Pack{{Size: 500}}

	packs := calculator.CalculatePacks(order)
	if !equalPacks(packs, expectedPacks) {
		t.Errorf("Expected packs %v, but got %v", expectedPacks, packs)
	} else {
		t.Log("TestPackCalculatorV1_CalculatePacks_251: Success")
	}
}

func TestPackCalculatorV1_CalculatePacks_501(t *testing.T) {
	order := models.Order{Items: 501}
	expectedPacks := []models.Pack{{Size: 500}, {Size: 250}}

	packs := calculator.CalculatePacks(order)
	if !equalPacks(packs, expectedPacks) {
		t.Errorf("Expected packs %v, but got %v", expectedPacks, packs)
	} else {
		t.Log("TestPackCalculatorV1_CalculatePacks_501: Success")
	}
}

func TestPackCalculatorV1_CalculatePacks_12001(t *testing.T) {
	order := models.Order{Items: 12001}
	expectedPacks := []models.Pack{{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250}}

	packs := calculator.CalculatePacks(order)
	if !equalPacks(packs, expectedPacks) {
		t.Errorf("Expected packs %v, but got %v", expectedPacks, packs)
	} else {
		t.Log("TestPackCalculatorV1_CalculatePacks_12001: Success")
	}
}

func equalPacks(a, b []models.Pack) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Size != b[i].Size {
			return false
		}
	}
	return true
}
