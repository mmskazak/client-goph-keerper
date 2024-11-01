package config

import (
	"dario.cat/mergo"
	"fmt"
)

// Config содержит поля конфигурации.
type Config struct {
	AppUrl string `json:"app_url"` // Путь до сервера goph-keeper
}

func NewConfig() *Config {
	return &Config{
		AppUrl: "app_url",
	}
}

// InitConfig инициализирует конфигурацию из флагов командной строки и переменных окружения.
func InitConfig() (*Config, error) {
	config := NewConfig()
	config, err := mergeFlags(config)
	if err != nil {
		return nil, fmt.Errorf("error merge: %w", err)
	}

	return config, nil
}

func mergeFlags(config *Config) (*Config, error) {
	configFromFlags := parseFlags()
	err := mergo.Merge(config, configFromFlags, mergo.WithOverride)
	if err != nil {
		return nil, fmt.Errorf("mergo from flags error: %v", err)
	}
	return config, nil
}
