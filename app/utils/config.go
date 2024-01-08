package config

import (
	config "github.com/miu200521358/mlib_go/pkg/utils"

)

type Config struct{}

// SaveConfig saves the config to the config file
func (c *Config) SaveConfig(key string, values []string, limit int) error {
	return config.SaveConfig(key, values, limit)
}

// LoadConfig loads the config from the config file
func (c *Config) LoadConfig(key string) []string {
	return config.LoadConfig(key)
}
