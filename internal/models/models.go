package models

// Order represents a customer order with a specific number of items.
type Order struct {
	// Items is the total number of items the customer wants to order.
	Items int `json:"items"`
}

// Pack represents a pack size with a specific number of items.
type Pack struct {
	// Size is the number of items contained in a single pack.
	Size int `json:"size"`
}