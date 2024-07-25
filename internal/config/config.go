package config

import (
    "encoding/json"
    "os"
    "github.com/Reaper1994/go-package-master/internal/models"
)

// Config represents the application configuration.
type Config struct {
    Packs []models.Pack `json:"packs"`
}

// LoadConfig loads configuration from a file.
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
