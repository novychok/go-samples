package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI      string
	DataBase string
}

func New(ctx context.Context, cfg *Config) (*mongo.Client, func(), error) {
	opts := options.Client().ApplyURI(cfg.URI)

	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, nil, err
	}

	cleaner := func() {
		_ = conn.Disconnect(ctx)
	}

	conn.Database(cfg.DataBase)

	return conn, cleaner, nil
}

func GetDatabase(conn *mongo.Client, cfg *Config) *mongo.Database {
	return conn.Database(cfg.DataBase)
}
