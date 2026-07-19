// Package config loads runtime configuration from the environment.
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Config holds RPC endpoints, chain selection and license keys.
type Config struct {
	RPCURL       string
	WSURL        string
	FlashbotsURL string
	ChainID      int64
	LicenseKeys  []string
}

// Load reads configuration from the process environment.
func Load() (*Config, error) {
	c := &Config{
		RPCURL:       os.Getenv("EVM_RPC_URL"),
		WSURL:        os.Getenv("EVM_WS_URL"),
		FlashbotsURL: envDefault("FLASHBOTS_URL", "https://relay.flashbots.net"),
		ChainID:      envInt("CHAIN_ID", 1),
	}
	if c.RPCURL == "" {
		return nil, errors.New("EVM_RPC_URL is required")
	}
	if keys := os.Getenv("LICENSE_KEYS"); keys != "" {
		c.LicenseKeys = strings.Split(keys, ",")
	}
	return c, nil
}

func envDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt(key string, fallback int64) int64 {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			return n
		}
	}
	return fallback
}
