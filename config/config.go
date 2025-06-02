package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	envEnvName  = "ENV"
	portEnvName = "PORT"
)

var ErrEmptyEnvParameter = errors.New("empty env parameter")

type Config struct {
	Env  string
	Port string
}

func New() (Config, error) {
	var c Config

	env := os.Getenv(envEnvName)
	if env == "" {
		return Config{}, fmt.Errorf("%w: %s", ErrEmptyEnvParameter, envEnvName)
	}
	c.Env = env

	port := os.Getenv(portEnvName)
	if port == "" {
		return Config{}, fmt.Errorf("%w: %s", ErrEmptyEnvParameter, portEnvName)
	}
	c.Port = port

	return c, nil
}
