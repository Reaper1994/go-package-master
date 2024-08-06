package services

import (
	"sort"

	"github.com/Reaper1994/go-package-master/internal/models"
)

type PackCalculatorV1 struct {
	Packs       []models.Pack // Original Packs
	SortedPacks []models.Pack // Cached Sorted Packs
}

// SortAndCachePacks sorts the Packs slice in descending order and stores the result.
func (pc *PackCalculatorV1) SortAndCachePacks() {
	// Copy the original packs to avoid modifying the original order
	pc.SortedPacks = make([]models.Pack, len(pc.Packs))
	copy(pc.SortedPacks, pc.Packs)

	// Sort the copied packs in descending order
	sort.Slice(pc.SortedPacks, func(i, j int) bool {
		return pc.SortedPacks[i].Size > pc.SortedPacks[j].Size
	})
}

// CalculatePacks calculates the optimal combination of packs.
func (pc *PackCalculatorV1) CalculatePacks(order models.Order) []models.Pack {
	// Ensure the packs are sorted and cached before using them
	pc.SortAndCachePacks()

	remaining := order.Items
	var res []models.Pack

	var orderLargePacks []models.Pack

	largestPackSize := pc.SortedPacks[0].Size
	quotient := remaining / largestPackSize
	remainder := remaining % largestPackSize

	// Add the required number of largest packs to the result
	for i := 0; i < quotient; i++ {
		orderLargePacks = append(orderLargePacks, models.Pack{Size: largestPackSize})
	}
	remaining = remainder

	// Packing the remaining items with the next available packs
	for remaining > 0 {
		packAdded := false

		// Try to use the largest pack that is less than or equal to the remaining items
		for _, pack := range pc.SortedPacks {
			if remaining >= pack.Size {
				res = append(res, pack)
				remaining -= pack.Size
				packAdded = true
				break
			}
		}

		// If no pack was added, use the smallest available pack to handle the remaining items
		if !packAdded {
			smallestPack := pc.SortedPacks[len(pc.SortedPacks)-1]
			res = append(res, smallestPack)
			remaining -= smallestPack.Size
		}
	}

	// Append the optimized smaller packs to the result
	result := append(orderLargePacks, optimizePacks(res, pc.SortedPacks)...)

	return result
}

// optimizePacks ensures that the final result has the minimal number of packs
func optimizePacks(packs []models.Pack, sortedPacks []models.Pack) []models.Pack {
	availablePackSizes := make(map[int]models.Pack)
	for _, pack := range sortedPacks {
		availablePackSizes[pack.Size] = pack
	}

	var consolidated []models.Pack
	for _, pack := range packs {
		consolidated = append(consolidated, pack)
	}

	for i := 0; i < len(consolidated)-1; i++ {
		for j := i + 1; j < len(consolidated); j++ {
			combinedSize := consolidated[i].Size + consolidated[j].Size
			if newPack, exists := availablePackSizes[combinedSize]; exists {
				consolidated[i] = newPack
				consolidated = append(consolidated[:j], consolidated[j+1:]...)
				break
			}
		}
	}

	return consolidated
}
