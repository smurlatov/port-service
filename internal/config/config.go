package config

import "os"

type Config struct {
	HttpAddr string
}

func Read() Config {
	var config Config
	addr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HttpAddr = addr
	}
	return config
}
