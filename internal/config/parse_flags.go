package config

import "flag"

func parseFlags() *Config {
	config := NewConfig()

	flag.StringVar((*string)(&config.LogLevel), "l", string(config.LogLevel), "Log level")
	flag.StringVar(&config.AppUrl, "app_url", "", "IP-адрес сервера goph-keeper")

	// Разбор командной строки
	flag.Parse()

	return config
}
