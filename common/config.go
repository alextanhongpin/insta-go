package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Port string `json:"port"`
}

const configPath string = "conf.json"

var configuration Configuration

// GetConfig reads the config from a json file
// and returns the config
func GetConfig() Configuration {
	// Check if the config is loaded or not
	if (Configuration{} != configuration) {
		return configuration
	}
	file, _ := os.Open(configPath)
	err := json.NewDecoder(file).Decode(&configuration)

	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	return configuration
}
