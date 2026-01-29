package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port                string   `json:"port"`
	Backends            []string `json:"backends"`
	HealthCheckInterval string   `json:"health_check_interval"`
}

func LoadConfig(file string) *Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal("Config file error: ", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		log.Fatal("Config decode error: ", err)
	}
	return &config
}
