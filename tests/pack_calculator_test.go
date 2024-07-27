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

	var calculatePacksFromRequestedItems func(int, []models.Pack) []models.Pack
	calculatePacksFromRequestedItems = func(itemsOrdered int, res []models.Pack) []models.Pack {
		if itemsOrdered == 0 {
			return res
		}

		for _, pack := range pc.Packs {
			if itemsOrdered >= pack.Size {
				res = append(res, pack)
				itemsOrdered -= pack.Size
				return calculatePacksFromRequestedItems(itemsOrdered, res)
			}
		}

		// If remaining items are less than the smallest pack size, add the smallest pack
		if itemsOrdered > 0 {
			smallestPack := pc.Packs[len(pc.Packs)-1]
			res = append(res, smallestPack)
		}

		return res
	}

	return calculatePacksFromRequestedItems(order.Items, []models.Pack{})
}
