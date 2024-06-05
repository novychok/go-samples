package psql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {

	psqlUri := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s host=%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"))

	db, err := sql.Open("postgres", psqlUri)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
