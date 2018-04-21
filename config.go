package main

import (
	"github.com/BurntSushi/toml"
)

// Target - target host to check
type Target struct {
	Host []string `toml:"hosts"`
}

// Refresh - how often should made check availability
type Refresh struct {
	Wait int `toml:"delay"`
}

// Config - struct for all data in config
type Config struct {
	Target  Target  `toml:"target"`
	Refresh Refresh `toml:"refresh"`
}

// NewConfig - create new Config instance
func NewConfig() *Config {
	return &Config{}
}

// Load - function load data from config
func (c *Config) Load(configPath string) error {
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		return err
	}

	return nil
}
