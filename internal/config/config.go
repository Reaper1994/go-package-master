package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Packs []struct {
		Size int `json:"size"`
	} `json:"packs"`
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(byteValue, &config)
	return config, err
}
