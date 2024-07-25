package test

import (
	"fmt"
	"testing"

	"github.com/Reaper1994/go-package-master/internal/models"
	"github.com/Reaper1994/go-package-master/internal/services"
)

func TestPackCalculatorV1_CalculatePacks(t *testing.T) {
	calculator := services.PackCalculatorV1{
		Packs: []models.Pack{
			{Size: 250},
			{Size: 500},
			{Size: 1000},
			{Size: 2000},
			{Size: 5000},
		},
	}

	tests := []struct {
		items         int
		expectedPacks []models.Pack
	}{
		{1, []models.Pack{{Size: 250}}},
		{250, []models.Pack{{Size: 250}}},
		{251, []models.Pack{{Size: 500}}},
		{501, []models.Pack{{Size: 500}, {Size: 250}}},
		{12001, []models.Pack{{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250}}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("items=%d", tt.items), func(t *testing.T) {
			order := models.Order{Items: tt.items}
			packs := calculator.CalculatePacks(order)
			if !equalPacks(packs, tt.expectedPacks) {
				t.Errorf("Expected packs %v, but got %v", tt.expectedPacks, packs)
			}
		})
	}
}

func equalPacks(a, b []models.Pack) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
