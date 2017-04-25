package common

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration is the model for the config
// loaded from config file
type Configuration struct {
	Port      string `json:"port"`
	JWTSecret string `json:"jwt_secret"`
}

const configPath string = "conf.json"

// Global Config variable
var Config Configuration

func init() {
	Config = loadConfig()
}

// LoadConfig returns the configuration that is loaded
// from conf.json
func loadConfig() Configuration {
	var config Configuration
	file, _ := os.Open(configPath)
	err := json.NewDecoder(file).Decode(&config)

	if err != nil {
		fmt.Printf("Error:%s", err)
	}
	return config
}
