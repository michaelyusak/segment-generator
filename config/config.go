package config

import (
	"encoding/json"
	"fmt"
	"michaelyusak/biaenergi-segment-generator.git/entity"
	"os"
)

type Neo4jConfig struct {
	Uri            string          `json:"uri"`
	Username       string          `json:"username"`
	Password       string          `json:"password"`
	ConnectTimeout entity.Duration `json:"connect_timeout"`
}

type ServiceConfig struct {
	Neo4j Neo4jConfig `json:"neo4j"`
}

type AppConfig struct {
	Port           string          `json:"port"`
	GracefulPeriod entity.Duration `json:"graceful_period"`
	Service        ServiceConfig   `json:"service"`
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
