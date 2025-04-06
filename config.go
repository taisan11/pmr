package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadConfig() (*Config, error) {
	// Define default config
	defaultConfig := Config{
		Level: []string{"scoop"},
	}

	// Try to open existing config file
	file, err := os.Open("config.json")
	if err != nil {
		// If file doesn't exist, create it with default config
		if os.IsNotExist(err) {
			return createDefaultConfig(defaultConfig)
		}
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Decode existing config
	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &config, nil
}

func createDefaultConfig(config Config) (*Config, error) {
	file, err := os.Create("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return nil, fmt.Errorf("failed to write config: %w", err)
	}

	return &config, nil
}

func loadPM(PMname string) (*PM, error) {
	// Try to open existing PM file
	file, err := os.Open(PMname + ".json")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Decode existing PM
	var pm PM
	if err := json.NewDecoder(file).Decode(&pm); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &pm, nil
}
