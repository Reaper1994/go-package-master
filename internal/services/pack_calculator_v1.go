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
	// Helper functions
	contains := func(packs []models.Pack, value int) bool {
		for _, pack := range packs {
			if pack.Size == value {
				return true
			}
		}
		return false
	}

	orderAsc := func(packs []models.Pack) []int {
		sizes := make([]int, len(packs))
		for i, pack := range packs {
			sizes[i] = pack.Size
		}
		sort.Ints(sizes)
		return sizes
	}

	findLowAndHighNearestToValue := func(packs []int, value int) (int, int) {
		low, high := -1, -1
		for _, pack := range packs {
			if pack <= value {
				low = pack
			}
			if pack >= value && high == -1 {
				high = pack
			}
		}
		return low, high
	}

	sum := func(packs []models.Pack) int {
		total := 0
		for _, pack := range packs {
			total += pack.Size
		}
		return total
	}

	getNearestValue := func(value int, packs []models.Pack) int {
		sizes := orderAsc(packs)
		_, high := findLowAndHighNearestToValue(sizes, value)
		return high
	}

	mapResults := func(packs []models.Pack) []models.Pack {
		return packs
	}

	var calculatePacksFromRequestedItems func(int, []models.Pack, []models.Pack) []models.Pack
	calculatePacksFromRequestedItems = func(itemsOrdered int, packs []models.Pack, res []models.Pack) []models.Pack {
		if contains(packs, itemsOrdered) {
			return mapResults(append(res, models.Pack{Size: itemsOrdered}))
		} else {
			remaining := itemsOrdered
			originalItemsOrdered := itemsOrdered

			if len(res) > 0 {
				temp := make([]int, len(packs))
				for i, pack := range packs {
					temp[i] = pack.Size
				}
				lowValue, highValue := findLowAndHighNearestToValue(orderAsc(packs), remaining)
				if remaining >= lowValue && remaining <= highValue {
					res = append(res, models.Pack{Size: highValue})
					remaining -= highValue
				} else if remaining >= lowValue {
					res = append(res, models.Pack{Size: lowValue})
					remaining -= lowValue
				}
			}

			if remaining > 0 {
				for _, pack := range packs {
					if remaining >= pack.Size {
						res = append(res, pack)
						remaining -= pack.Size
						return calculatePacksFromRequestedItems(remaining, packs, res)
					}
				}
				var smallestPack = packs[len(packs)-1]
				if remaining < smallestPack.Size {
					res = append(res, smallestPack)
				}
			}

			// Perform validation to ensure we are sending out the smallest number of packs
			sumResults := sum(res)
			if contains(packs, sumResults) {
				res = []models.Pack{{Size: sumResults}}
			} else if nearest := getNearestValue(originalItemsOrdered, packs); nearest < sumResults && nearest > originalItemsOrdered {
				res = []models.Pack{{Size: nearest}}
			}
		}
		return mapResults(res)
	}

	return calculatePacksFromRequestedItems(order.Items, pc.Packs, []models.Pack{})
}
