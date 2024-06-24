package config

import (
	"os"

	"github.com/novychok/go-samples/realtime/internal/pkg/postgres"
)

type Config struct {
	PostgresConfig *postgres.Config
}

func New() *Config {
	return &Config{
		PostgresConfig: &postgres.Config{
			Name:     os.Getenv("POSTGRES_DB"),
			Host:     os.Getenv("REALTIME_PGHOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Sslmode:  os.Getenv("REALTIME_PGSSLMODE"),
		},
	}
}
