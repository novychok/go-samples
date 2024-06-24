package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Name     string
	Host     string
	User     string
	Password string
	Sslmode  string
}

type Connection *sql.DB

func New(cfg *Config) (Connection, func(), error) {

	uri := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s",
		cfg.User, cfg.Name, cfg.Sslmode, cfg.Password, cfg.Host)

	fmt.Println(uri)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	cleaner := func() {
		_ = db.Close()
	}

	return db, cleaner, nil
}
