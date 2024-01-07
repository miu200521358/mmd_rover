package app

import (
	"github.com/miu200521358/mlib_go/utils"
)

// SaveConfig saves the config to the config file
func (a *App) SaveConfig(key string, values []string, limit int) error {
	return utils.SaveConfig(key, values, limit)
}

// LoadConfig loads the config from the config file
func (a *App) LoadConfig(key string) []string {
	return utils.LoadConfig(key)
}
