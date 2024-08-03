package transformers

import (
	"encoding/json"

	"github.com/Reaper1994/go-package-master/internal/models"
)

func FormatPacks(packs []models.Pack) string {
	counts := make(map[int]int)

	// Count the occurrences of each pack size
	for _, pack := range packs {
		counts[pack.Size]++
	}

	// Convert the map to JSON
	result, err := json.Marshal(counts)
	if err != nil {
		// Handle the error appropriately; for now, return an empty JSON object
		return "{}"
	}

	return string(result)
}
