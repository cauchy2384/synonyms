package main

import (
	"os"
	"strconv"
)

type Config struct {
	Address                string
	ShutdownTimeoutSeconds uint64
}

func LoadConfig() Config {
	const (
		defaultServerAddress       = ":1337"
		defaultShutdownTimeSeconds = 10
	)

	return Config{
		Address:                getEnvOrDefault("SERVER_ADDRESS", defaultServerAddress),
		ShutdownTimeoutSeconds: getEnvOrDefaultUint64("SHUTDOWN_TIMEOUT_SECONDS", defaultShutdownTimeSeconds),
	}
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvOrDefaultUint64(key string, defaultValue uint64) uint64 {
	const (
		base    = 10
		bitSize = 64
	)

	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseUint(value, base, bitSize); err == nil {
			return v
		}
	}
	return defaultValue
}
