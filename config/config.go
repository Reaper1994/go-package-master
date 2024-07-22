package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Packs []Pack `json:"packs"`
}

type Pack struct {
	Size int `json:"size"`
}

// LoadConfig loads the configuration from a JSON file located at filepath.
// It returns a Config struct and an error if any issues are encountered.
func LoadConfig(filepath string) (Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
