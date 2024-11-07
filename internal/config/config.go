package config

import (
	"net"

	"github.com/joho/godotenv"
)

// HTTPConfigSearcher interface for search HTTP config.
type HTTPConfigSearcher interface {
	Get() (*HTTPConfig, error)
}

// Load dotenv from path to env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

// HTTPConfig is config for HTTP
type HTTPConfig struct {
	Host string
	Port string
}

// Address get address from config
func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
