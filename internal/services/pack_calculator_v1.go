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
// Cache in the sense refers to storing the sorted version of the Packs slice in a separate field (SortedPacks)
// within the PackCalculatorV1 struct
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

	// Optimize the result to minimize the number of packs used
	return optimizePacks(res, pc.SortedPacks)
}

// optimizePacks ensures that the final result has the minimal number of packs
func optimizePacks(packs []models.Pack, sortedPacks []models.Pack) []models.Pack {

	// Create a map of pack sizes to quickly look up if a combined size exists
	availablePackSizes := make(map[int]models.Pack)
	for _, pack := range sortedPacks {
		availablePackSizes[pack.Size] = pack
	}

	var consolidated []models.Pack
	for i := 0; i < len(packs); i++ {
		consolidated = append(consolidated, packs[i])
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
