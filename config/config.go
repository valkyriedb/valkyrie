package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	envEnvName  = "ENV"
	portEnvName = "PORT"
	passEnvName = "PASSWORD"
)

var ErrEmptyEnvParameter = errors.New("empty env parameter")

type Config struct {
	Env      string
	Port     string
	Password string
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

	pass := os.Getenv(passEnvName)
	if port == "" {
		return Config{}, fmt.Errorf("%w: %s", ErrEmptyEnvParameter, portEnvName)
	}
	c.Password = pass

	return c, nil
}
