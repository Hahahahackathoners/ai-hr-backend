package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig GRPCConfig `yaml:"server" validate:"required"`
}

type GRPCConfig struct {
	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}
	return nil
}
