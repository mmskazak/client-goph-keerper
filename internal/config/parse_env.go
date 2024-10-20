package config

import "os"

func parseEnv() *Config {
	config := NewConfig()

	if envLogLevel, ok := os.LookupEnv("LOG_LEVEL"); ok {
		config.LogLevel = LogLevel(envLogLevel)
	}

	if appUrl, ok := os.LookupEnv("APP_URL"); ok {
		config.AppUrl = appUrl
	}

	return config
}
