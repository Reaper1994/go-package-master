package services

import (
	"sort"
	"sync"

	"github.com/Reaper1994/go-package-master/internal/models"
)

type PackCalculatorV1 struct {
	Packs       []models.Pack // Original Packs
	SortedPacks []models.Pack // Cached Sorted Packs
	mu          sync.Mutex    // Mutex to protect SortedPacks
}

// SortAndCachePacks sorts the Packs slice in descending order and stores the result.
func (pc *PackCalculatorV1) SortAndCachePacks() {
	pc.mu.Lock()
	defer pc.mu.Unlock()

	// Copy the original packs to avoid modifying the original order
	pc.SortedPacks = make([]models.Pack, len(pc.Packs))
	copy(pc.SortedPacks, pc.Packs)

	// Perform the sorting concurrently
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		sort.Slice(pc.SortedPacks, func(i, j int) bool {
			return pc.SortedPacks[i].Size > pc.SortedPacks[j].Size
		})
	}()

	wg.Wait()
}

// CalculatePacks calculates the optimal combination of packs.
func (pc *PackCalculatorV1) CalculatePacks(order models.Order) []models.Pack {
	// Ensure the packs are sorted and cached before using them
	pc.SortAndCachePacks()

	remaining := order.Items
	var res []models.Pack
	var orderLargePacks []models.Pack

	var wg sync.WaitGroup
	largestPackSize := pc.SortedPacks[0].Size
	quotient := remaining / largestPackSize
	remainder := remaining % largestPackSize

	// Calculate the largest packs concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < quotient; i++ {
			orderLargePacks = append(orderLargePacks, models.Pack{Size: largestPackSize})
		}
	}()

	// Calculate remaining packs concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		tempRemaining := remainder
		for tempRemaining > 0 {
			packAdded := false

			// Try to use the largest pack that is less than or equal to the remaining items
			for _, pack := range pc.SortedPacks {
				if tempRemaining >= pack.Size {
					pc.mu.Lock()
					res = append(res, pack)
					pc.mu.Unlock()
					tempRemaining -= pack.Size
					packAdded = true
					break
				}
			}

			// If no pack was added, use the smallest available pack to handle the remaining items
			if !packAdded {
				smallestPack := pc.SortedPacks[len(pc.SortedPacks)-1]
				pc.mu.Lock()
				res = append(res, smallestPack)
				pc.mu.Unlock()
				tempRemaining -= smallestPack.Size
			}
		}
	}()

	wg.Wait()

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