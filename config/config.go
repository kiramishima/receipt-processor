package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/kiramishima/receipt-processor/domain"
	"go.uber.org/fx"
)

// Load config from enviroment
func Load() (*domain.Configuration, error) {
	var cfg domain.Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// NewConfig creates and load config
func NewConfig() *domain.Configuration {
	cfg, err := Load()
	if err != nil {
		log.Printf("Can't load the configuration. Error: %s", err.Error())
	}

	return cfg
}

// Module
var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(NewConfig),
)
