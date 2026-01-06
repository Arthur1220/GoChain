package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	RPCURL      string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/gochain?sslmode=disable"),
		RPCURL:      getEnv("RPC_URL", "https://eth.llamarpc.com"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
