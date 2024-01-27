package config

import (
	"embed"
	"encoding/json"
	"log"
	"os"
)

// Appshell config structure as a map
type Config struct {
	Appshells map[string]map[string]string `json:"appshells"`
}

//go:embed appshell-config.json
var configFS embed.FS

// Load the embedded appshell config
func LoadAppShellConfig() (Config, error) {
	// Logger that writes to STDOUT
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	var config Config
	data, err := configFS.ReadFile("appshell-config.json")
	if err != nil {
		logger.Println("ERROR reading appshell config")
		return config, err
	}
	err = json.Unmarshal(data, &config)
	return config, err
}
