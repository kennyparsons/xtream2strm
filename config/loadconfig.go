package config

import (
	"os"
	"xtream2strm/models"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configPath string) (models.Config, error) {
	var config models.Config

	// Open the configuration file
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Decode the YAML configuration file into the Config struct
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
