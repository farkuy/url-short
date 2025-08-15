package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ENV          string `yaml:"env"`
	STORAGE_PATH string `yaml:"storage_path"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	ADDRESS      string        `yaml:"address"`
	TIMEOUT      time.Duration `yaml:"timeout"`
	IDLE_TIMEOUT time.Duration `yaml:"idle_timeout"`
}

func LoadConfig() (*Config, error) {
	localYamlPath := os.Getenv("LOCAL_YAML")
	if localYamlPath == "" {
		return nil, errors.New("Path for local.yaml file not found")
	}

	data, err := os.ReadFile(localYamlPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
