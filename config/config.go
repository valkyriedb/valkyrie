package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	portEnvName = "PORT"
)

var ErrEmptyEnvParameter = errors.New("empty env parameter")

type Config struct {
	Port string
}

func New() (Config, error) {
	var c Config

	port := os.Getenv(portEnvName)
	if port == "" {
		return Config{}, fmt.Errorf("%w: %s", ErrEmptyEnvParameter, portEnvName)
	}
	c.Port = port

	return c, nil
}
