package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Reaper1994/go-package-master/models"
)

func TestCalculateHandlerV1(t *testing.T) {
	packs := []models.Pack{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}

	handler := &CalculateHandlerV1{Packs: packs}

	tests := []struct {
		order       models.Order
		expected    []models.Pack
		description string
	}{
		{
			order:       models.Order{Items: 1},
			expected:    []models.Pack{{Size: 250}},
			description: "1 item",
		},
		{
			order:       models.Order{Items: 250},
			expected:    []models.Pack{{Size: 250}},
			description: "250 items",
		},
		{
			order:       models.Order{Items: 251},
			expected:    []models.Pack{{Size: 500}},
			description: "251 items",
		},
		{
			order:       models.Order{Items: 501},
			expected:    []models.Pack{{Size: 500}, {Size: 250}},
			description: "501 items",
		},
		{
			order:       models.Order{Items: 12001},
			expected:    []models.Pack{{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250}},
			description: "12001 items",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			body, err := json.Marshal(test.order)
			if err != nil {
				t.Fatalf("failed to marshal order: %v", err)
			}

			req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			var result []models.Pack
			err = json.NewDecoder(rr.Body).Decode(&result)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if len(result) != len(test.expected) {
				t.Fatalf("expected %v packs, got %v packs", len(test.expected), len(result))
			}
			for i := range result {
				if result[i].Size != test.expected[i].Size {
					t.Fatalf("expected pack size %v, got pack size %v", test.expected[i].Size, result[i].Size)
				}
			}
		})
	}
}

func TestCalculateHandlerV2(t *testing.T) {
	packs := []models.Pack{
		{Size: 250},
		{Size: 500},
		{Size: 1000},
		{Size: 2000},
		{Size: 5000},
	}

	handler := &CalculateHandlerV2{Packs: packs}

	tests := []struct {
		order       models.Order
		expected    []models.Pack
		description string
	}{
		{
			order:       models.Order{Items: 1},
			expected:    []models.Pack{{Size: 250}},
			description: "1 item",
		},
		{
			order:       models.Order{Items: 250},
			expected:    []models.Pack{{Size: 250}},
			description: "250 items",
		},
		{
			order:       models.Order{Items: 251},
			expected:    []models.Pack{{Size: 500}},
			description: "251 items",
		},
		{
			order:       models.Order{Items: 501},
			expected:    []models.Pack{{Size: 500}, {Size: 250}},
			description: "501 items",
		},
		{
			order:       models.Order{Items: 12001},
			expected:    []models.Pack{{Size: 5000}, {Size: 5000}, {Size: 2000}, {Size: 250}},
			description: "12001 items",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			body, err := json.Marshal(test.order)
			if err != nil {
				t.Fatalf("failed to marshal order: %v", err)
			}

			req, err := http.NewRequest("POST", "/api/v2/calculate", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Fatalf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			var result []models.Pack
			err = json.NewDecoder(rr.Body).Decode(&result)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if len(result) != len(test.expected) {
				t.Fatalf("expected %v packs, got %v packs", len(test.expected), len(result))
			}
			for i := range result {
				if result[i].Size != test.expected[i].Size {
					t.Fatalf("expected pack size %v, got pack size %v", test.expected[i].Size, result[i].Size)
				}
			}
		})
	}
}
