package services

import (
	"sort"

	"github.com/Reaper1994/go-package-master/internal/models"
)

type PackCalculatorV1 struct {
	Packs []models.Pack
}

// CalculatePacks calculates the optimal combination of packs.
func (pc *PackCalculatorV1) CalculatePacks(order models.Order) []models.Pack {
	// Sort packs in descending order to prioritize larger packs first
	sort.Slice(pc.Packs, func(i, j int) bool {
		return pc.Packs[i].Size > pc.Packs[j].Size
	})

	remaining := order.Items
	var res []models.Pack

	for remaining > 0 {
		packAdded := false

		// Try to use the largest pack that is less than or equal to the remaining items
		for _, pack := range pc.Packs {
			if remaining >= pack.Size {
				res = append(res, pack)
				remaining -= pack.Size
				packAdded = true
				break
			}
		}
		if packAdded {
			continue
		}

		// If no pack was added, use the smallest available pack to handle the remaining items
		if remaining > 0 {
			smallestPack := pc.Packs[len(pc.Packs)-1]
			res = append(res, smallestPack)
			remaining -= smallestPack.Size
		}
	}

	// Optimization: Consolidate results to avoid unnecessary packs
	res = optimizePacks(res, pc.Packs)

	return res
}

// optimizePacks ensures that the final result has the minimal number of packs
func optimizePacks(packs []models.Pack, availablePacks []models.Pack) []models.Pack {
	sort.Slice(packs, func(i, j int) bool {
		return packs[i].Size > packs[j].Size
	})

	for i := 0; i < len(packs); i++ {
		for j := i + 1; j < len(packs); j++ {
			if packs[i].Size == packs[j].Size {
				continue
			}
			if packs[i].Size > packs[j].Size {
				remaining := packs[i].Size - packs[j].Size
				for _, pack := range availablePacks {
					if remaining == pack.Size {
						packs[i] = pack
						packs = append(packs[:j], packs[j+1:]...)
						i = -1
						break
					}
				}
			}
		}
	}

	return packs
}
