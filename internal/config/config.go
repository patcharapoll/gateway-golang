package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(NewConfiguration)

// Configuration ...
type Configuration struct {
	Port        string `env:"PORT" envDefault:"3003"`
	Environment string `env:"ENV" envDefault:"localhost"`

	EXAMPLEConnection string        `env:"EXAMPLE_CONNECTION" envDefault:"localhost:3001"`
	TokenDuration     time.Duration `env:"TOKEN_DURATION" envDefault:"30m"`
}

// NewConfiguration ...
func NewConfiguration() (*Configuration, error) {
	cfg := new(Configuration)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
