package config

import (
	"os"

	"github.com/novychok/go-samples/mongorepo/internal/pkg/mongo"
)

type Config struct {
	Mongo *mongo.Config
}

func New() *Config {
	return &Config{
		Mongo: &mongo.Config{
			URI:      "mongodb://127.0.0.1:27017",
			DataBase: os.Getenv("DATABASE_NAME"),
		},
	}
}
