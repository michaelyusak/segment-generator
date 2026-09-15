package config

import (
	"encoding/json"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"os"
)

type AppConfig struct {
	Port           string          `json:"port"`
	GracefulPeriod entity.Duration `json:"graceful_period"`
}

func Init() (*AppConfig, error) {
	filePath := os.Getenv("SEGMENT_GENERATOR_CONFIG_JSON")
	if filePath == "" {
		return nil, fmt.Errorf("[config][Init] config file path not set")
	}

	configData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("[config][Init] failed to open config file: %w", err)
	}

	var config AppConfig

	err = json.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("[config][Init] failed to unmarshal: %w", err)
	}

	return &config, nil
}
