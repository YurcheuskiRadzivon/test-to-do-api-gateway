package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const (
	StorageMinio      = "MINIO"
	StorageLocalstack = "LOCALSTACK"
	StorageFS         = "FS"
	DefaultStorage    = "FS"
)

type (
	Config struct {
		HTTP HTTP
		APP  APP
	}

	HTTP struct {
		PORT string `env:"HTTP_PORT,required"`
	}

	APP struct {
		DOMAIN string `env:"APP_DOMAIN,required"`
	}
)

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
