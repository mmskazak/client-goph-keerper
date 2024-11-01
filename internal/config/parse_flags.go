package config

import "flag"

func parseFlags() *Config {
	config := NewConfig()
	flag.StringVar(&config.AppUrl, "app_url", "", "IP-адрес сервера goph-keeper")
	// Разбор командной строки
	flag.Parse()
	return config
}
